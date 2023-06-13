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

	"github.com/shopspring/decimal"
)

func (os *OrderbooksService) ProcessMarketOrder(ctx context.Context, userID string, pair *Pair.Pair, side orderbook.SideOrder, amount string) (*OrderResult, error) {

	// 1. check if user can pay
	userWallet, err := wallet.GetWallet(ctx, userID, []string{
		pair.Base,
		pair.Quote,
	})

	if err != nil {
		return nil, err
	}
	fmt.Printf("wallet: %v\n", userWallet)
	fmt.Printf("pair: %v\n", pair)

	var currency string

	if side == orderbook.SideOrder_BUY {
		currency = pair.Quote
	} else {
		currency = pair.Base
	}

	quantity, err := decimal.NewFromString(amount)
	if err != nil {
		return nil, err
	}

	quantityInt := price_helper.ToIntWithAppCoef(&quantity)

	fmt.Printf("currency: %v\n", currency)
	enough := userWallet.HasEnough(currency, &quantityInt)

	if !enough {
		return nil, errors.New("insufficient funds")
	}
	// fmt.Printf("os.OrderbookClients: %v\n", os.OrderbookClients)

	if os.HasPair(pair.Symbol()) {
		os.Start(pair.Symbol())
	}

	// 2. send message to orderbook
	orderResponse, err := os.OrderbookClients[pair.Symbol()].Client.NewMarketOrder(ctx, &orderbook.NewMarketOrderArgs{
		Side:     side,
		Quantity: price_helper.ToIntWithAppCoef(&quantity),
		Pair:     pair.Symbol(),
	})

	if err != nil {
		fmt.Printf("NewMarketOrder err: %v\n", err)
		return nil, err
	}

	fmt.Printf("orderResponse: %v\n", orderResponse)

	paid, quantityBought, averagePrice := ProcessFilledOrders(ProcessFilledOrdersArgs{
		Orders:                   orderResponse.Order,
		Partial:                  orderResponse.Partial,
		PartialQuantityProcessed: &orderResponse.PartialQuantityProcessed,
		UserId:                   userID,
		Pair:                     pair,
	})
	fmt.Printf("quantityBought: %v\n", quantityBought)
	fmt.Printf("averagePrice: %v\n", averagePrice)
	fmt.Printf("paid: %v\n", paid)

	orderEntity := order_entity.NewOrder(order_entity.Order{
		UserID:           userID,
		Symbol:           pair.Symbol(),
		Side:             constants.SideOrder(side),
		Price:            averagePrice,
		FilledQuantity:   quantityBought,
		Paid:             paid,
		OriginalQuantity: quantityInt,
		Type:             constants.Order_MARKET,
	})
	fmt.Printf("quantityBought: %v\n", quantityBought)
	fmt.Printf("quantity: %v\n", quantity)

	if quantityBought == 0 {
		orderEntity.Status = constants.Order_EMPTY
	} else if quantityBought == quantityInt {
		orderEntity.Status = constants.Order_FILLED
	} else {
		orderEntity.Status = constants.Order_PARTIALLY_FILLED
	}

	go wallet.UpdateWallet(wallet.UpdateWalletArgs{
		Ctx:    context.TODO(),
		UserID: userID,
		Currencies: map[string]wallet.UpdateWalletCurrency{
			pair.Base: {
				Own: quantityBought,
			},
			pair.Quote: {
				Own: -paid,
			},
		},
		Wallet: userWallet,
	})

	go SaveOrder(orderEntity, &quantityBought)

	var status constants.StatusOrder = constants.Order_EMPTY
	if paid > 0 {
		status = constants.Order_PARTIALLY_FILLED

		if quantityBought == quantityInt {
			status = constants.Order_FILLED
		}

		// payment event
		filledOrderEvent := event_entity.NewPaymentEvent(&event_entity.PaymentEvent{
			UserID:              userID,
			Type:                event_entity.PayementEventType_MarketOrder,
			AmountAdded:         quantityBought,
			CurrencyAdded:       pair.Base,
			CurrencySubstracted: pair.Quote,
			AmountSubstracted:   paid,
		})

		event.Add(context.TODO(), filledOrderEvent)

	}

	transactionTime, err := time.Parse(time.RFC3339, orderResponse.GetTransactionTime())

	if err != nil {
		return nil, err
	}

	return &OrderResult{
		Symbol:                  pair.Symbol(),
		OrderId:                 orderEntity.ID,
		CumulativeQuoteQuantity: &paid,
		OriginalQuantity:        &quantityInt,
		ExecutedQuantity:        &quantityBought,
		Status:                  status,
		Type:                    constants.Order_LIMIT,
		Side:                    side,
		Price:                   &averagePrice,
		TransactionTime:         &transactionTime,
	}, nil
}
