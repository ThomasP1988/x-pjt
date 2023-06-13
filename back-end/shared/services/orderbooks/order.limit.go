package orderbooks

import (
	Pair "NFTM/shared/components/pair"
	"NFTM/shared/constants"
	event_entity "NFTM/shared/entities/event"
	order_entity "NFTM/shared/entities/order"
	price_helper "NFTM/shared/libs/price"
	orderbook "NFTM/shared/orderbook/grpc"
	"NFTM/shared/repositories/event"
	"NFTM/shared/repositories/wallet"
	"context"
	"errors"
	"fmt"
	"time"
)

func (os *OrderbooksService) ProcessLimitOrder(ctx context.Context, userID string, pair *Pair.Pair, side orderbook.SideOrder, amount string, price string) (*OrderResult, error) {
	fmt.Printf("\"ProcessLimitOrder\": %v\n", pair.Symbol())
	// 1. check if user can pay
	userWallet, err := wallet.GetWallet(ctx, userID, []string{
		pair.Base,
		pair.Quote,
	})
	fmt.Printf("amount: %v\n", amount)
	fmt.Printf("price: %v\n", price)
	if err != nil {
		return nil, err
	}

	quantityInt, err := price_helper.StringToIntWithAppCoef(amount)

	if err != nil {
		return nil, err
	}

	priceInt, err := price_helper.StringToIntWithAppCoef(price)

	if err != nil {
		return nil, err
	}

	var enough bool
	var currencyIn string
	var currencyOut string
	var availabilityAmount *int64

	if side == orderbook.SideOrder_BUY {
		currencyOut = pair.Quote
		currencyIn = pair.Base
		availabilityAmount = &priceInt
	} else {
		currencyOut = pair.Base
		currencyIn = pair.Quote
		availabilityAmount = &quantityInt
	}

	enough = userWallet.HasEnoughAvailability(currencyOut, availabilityAmount)

	if !enough {
		return nil, errors.New("insufficient funds")
	}

	fmt.Printf("quantityInt: %v\n", quantityInt)
	fmt.Printf("priceInt: %v\n", priceInt)
	orderEntity := order_entity.NewOrder(order_entity.Order{
		UserID:           userID,
		Symbol:           pair.Symbol(),
		Side:             constants.SideOrder(side),
		Price:            priceInt,
		Quantity:         quantityInt,
		Type:             constants.Order_LIMIT,
		OriginalQuantity: quantityInt,
	})

	if os.HasPair(pair.Symbol()) {
		os.Start(pair.Symbol())
	}

	// 2. send message to orderbook
	orderResponse, err := os.OrderbookClients[pair.Symbol()].Client.NewLimitOrder(ctx, &orderbook.NewLimitOrderArgs{
		Side:     side,
		Quantity: quantityInt,
		Price:    priceInt,
		Pair:     pair.Symbol(),
		OrderId:  orderEntity.ID,
	})

	if err != nil {
		return nil, err
	}

	fmt.Printf("orderResponse: %v\n", orderResponse)

	// paid, quantityBought := ProcessFilledOrders(orderResponse.Order, , &orderResponse.PartialQuantityProcessed, userID)
	paid, quantityBought, averagePrice := ProcessFilledOrders(ProcessFilledOrdersArgs{
		Orders:                   orderResponse.Order,
		Partial:                  orderResponse.Partial,
		PartialQuantityProcessed: &orderResponse.PartialQuantityProcessed,
		UserId:                   userID,
		Pair:                     pair,
		Side:                     constants.SideOrder(side),
	})

	fmt.Printf("quantityBought: %v\n", quantityBought)
	fmt.Printf("averagePrice: %v\n", averagePrice)
	fmt.Printf("paid: %v\n", paid)

	orderEntity.AveragePrice = averagePrice
	orderEntity.Paid = paid

	go SaveOrder(orderEntity, &quantityBought)

	// TODO update wallet taker

	fmt.Printf("wallet sub (quantityInt - quantityBought): %v\n", (quantityInt - quantityBought))

	go wallet.UpdateWallet(wallet.UpdateWalletArgs{
		Ctx:    context.TODO(),
		UserID: userID,
		Currencies: map[string]wallet.UpdateWalletCurrency{
			currencyIn: {
				Own: quantityBought,
			},
			currencyOut: {
				Own:       -paid,
				Available: -*availabilityAmount,
			},
		},
	})

	var status constants.StatusOrder = constants.Order_OPEN

	if paid > 0 {
		// payment event
		status = constants.Order_PARTIALLY_FILLED
		var typePayment event_entity.TypePaymentEvent
		if quantityBought == quantityInt {
			typePayment = event_entity.PayementEventType_LimitOrder
		} else {
			typePayment = event_entity.PayementEventType_LimitOrderPartial
		}

		filledOrderEvent := event_entity.NewPaymentEvent(&event_entity.PaymentEvent{
			UserID:              userID,
			Type:                typePayment,
			AmountAdded:         quantityBought,
			CurrencyAdded:       pair.Base,
			CurrencySubstracted: pair.Quote,
			AmountSubstracted:   paid,
		})

		go event.Add(context.TODO(), filledOrderEvent)
	}

	transactionTime, err := time.Parse(time.RFC3339, orderResponse.GetTransactionTime())

	if err != nil {
		return nil, err
	}

	return &OrderResult{
		Symbol:                  pair.Symbol(),
		OrderId:                 orderEntity.ID,
		TransactionTime:         &transactionTime,
		CumulativeQuoteQuantity: &paid,
		OriginalQuantity:        &quantityInt,
		ExecutedQuantity:        &quantityBought,
		Status:                  status,
		Type:                    constants.Order_LIMIT,
		Side:                    side,
		Price:                   &averagePrice,
	}, nil
}
