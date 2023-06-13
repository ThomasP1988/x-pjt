package server

import (
	price_helper "NFTM/shared/libs/price"
	grpc_ob "NFTM/shared/orderbook/grpc"
	"context"
	"fmt"
	"orderbook_service/lib/hub"
	"time"
)

func (s *OrderbookServer) NewMarketOrder(ctx context.Context, args *grpc_ob.NewMarketOrderArgs) (*grpc_ob.NewMarketOrderResponse, error) {

	orderbook := hub.GetOderbookByPair(args.GetPair())

	quantity := price_helper.FromIntWithAppCoef(args.Quantity)

	done, partial, partialQuantityProcessed, quantityLeft, err := orderbook.ProcessMarketOrder(FormatSide(args.Side), quantity)
	transactionTime := time.Now()
	if err != nil {
		return nil, err
	}

	fmt.Printf("done: %v\n", done)

	doneOrders := []*grpc_ob.Order{}

	for _, item := range done {
		doneOrders = append(doneOrders, ParseOrder(item))
	}

	response := &grpc_ob.NewMarketOrderResponse{
		Order:                    doneOrders,
		PartialQuantityProcessed: price_helper.ToIntWithAppCoef(&partialQuantityProcessed),
		QuantityLeft:             price_helper.ToIntWithAppCoef(&quantityLeft),
		TransactionTime:          transactionTime.Format(time.RFC3339),
	}

	if partial != nil {
		response.Partial = ParseOrder(partial)
	}

	return response, nil
}
