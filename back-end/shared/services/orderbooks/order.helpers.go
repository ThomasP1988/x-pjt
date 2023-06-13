package orderbooks

import (
	Pair "NFTM/shared/components/pair"
	"NFTM/shared/constants"
	event_entity "NFTM/shared/entities/event"
	order_entity "NFTM/shared/entities/order"
	price_helper "NFTM/shared/libs/price"
	orderbook "NFTM/shared/orderbook/grpc"
	"NFTM/shared/repositories/event"
	"NFTM/shared/repositories/order"
	"NFTM/shared/repositories/wallet"
	"context"
	"fmt"
)

type ProcessFilledOrdersArgs struct {
	Orders                   []*orderbook.Order
	Partial                  *orderbook.Order
	PartialQuantityProcessed *int64
	UserId                   string
	Pair                     *Pair.Pair
	Side                     constants.SideOrder
}

func ProcessFilledOrders(args ProcessFilledOrdersArgs) (int64, int64, int64) {

	var orderNumber int = len(args.Orders)

	var paid int64 = 0
	var quantityBought int64 = 0
	var totalPrice int64 = 0
	var averagePrice int64 = 0

	for _, orderResponse := range args.Orders {
		totalPrice += orderResponse.Price
		quantityBought += orderResponse.Quantity
		paid += orderResponse.Price * orderResponse.Quantity / price_helper.MultiplyBy

		go HandleFilledOrder(HandleFilledOrderArgs{
			Price:       &orderResponse.Price,
			Quantity:    &orderResponse.Quantity,
			OrderFilled: orderResponse,
			IsPartial:   false,
			UserId:      args.UserId,
			Pair:        args.Pair,
			Side:        args.Side,
		})
	}

	if args.Partial != nil {
		totalPrice += args.Partial.Price
		quantityBought += args.Partial.Quantity
		paid += args.Partial.Price * args.Partial.Quantity / price_helper.MultiplyBy

		go HandleFilledOrder(HandleFilledOrderArgs{
			Price:                    &args.Partial.Price,
			Quantity:                 &args.Partial.Quantity,
			OrderFilled:              args.Partial,
			IsPartial:                true,
			PartialQuantityProcessed: args.PartialQuantityProcessed,
			UserId:                   args.UserId,
			Pair:                     args.Pair,
			Side:                     args.Side,
		})
		orderNumber++
	}

	if orderNumber > 0 {
		averagePrice = totalPrice / int64(orderNumber)
	}

	fmt.Printf("quantityBought: %v\n", quantityBought)

	return paid, quantityBought, averagePrice
}

func SaveOrder(orderEntity *order_entity.Order, quantityBought *int64) {
	// TODO: pas bon
	orderEntity.FilledQuantity = *quantityBought
	order.Create(context.Background(), orderEntity)
}

type HandleFilledOrderArgs struct {
	Price                    *int64
	Quantity                 *int64
	OrderFilled              *orderbook.Order
	IsPartial                bool
	PartialQuantityProcessed *int64
	UserId                   string
	Pair                     *Pair.Pair
	Side                     constants.SideOrder
}

func HandleFilledOrder(args HandleFilledOrderArgs) {
	// UPDATE ORDERS
	if args.IsPartial {
		go order.PartiallyFilled(context.Background(), args.OrderFilled.Id, *args.PartialQuantityProcessed)
	} else {
		go order.Filled(context.Background(), args.OrderFilled.Id)
	}

	// SAVE payment event

	var currencyAdded string
	var amountAdded int64
	var currencySubstracted string
	var amountSubstracted int64
	var paymentEventType event_entity.TypePaymentEvent

	if args.IsPartial {
		paymentEventType = event_entity.PayementEventType_LimitOrder
	} else {
		paymentEventType = event_entity.PayementEventType_LimitOrderPartial
	}

	if args.Side == constants.Order_BUY {
		amountAdded = *args.Quantity
		currencyAdded = args.Pair.Base
		amountSubstracted = *args.Price * *args.Quantity / price_helper.MultiplyBy
		currencySubstracted = args.Pair.Quote
	}

	if args.Side == constants.Order_SELL {
		amountSubstracted = *args.Quantity
		currencySubstracted = args.Pair.Base
		amountAdded = *args.Price * *args.Quantity / price_helper.MultiplyBy
		currencyAdded = args.Pair.Quote
	}

	filledOrderEvent := event_entity.NewPaymentEvent(&event_entity.PaymentEvent{
		UserID:              args.UserId,
		Type:                paymentEventType,
		AmountAdded:         amountAdded,
		CurrencyAdded:       currencyAdded,
		CurrencySubstracted: currencySubstracted,
		AmountSubstracted:   amountSubstracted,
	})

	event.Add(context.TODO(), filledOrderEvent)

	// UPDATE maker wallet
	wallet.UpdateWallet(wallet.UpdateWalletArgs{
		Ctx:    context.TODO(),
		UserID: args.UserId,
		Currencies: map[string]wallet.UpdateWalletCurrency{
			currencyAdded: {
				Own:       amountAdded,
				Available: amountAdded,
			},
			currencySubstracted: {
				Own:       -amountSubstracted,
				Available: -amountSubstracted,
			},
		},
	})

}
