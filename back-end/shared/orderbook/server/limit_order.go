package server

import (
	price_helper "NFTM/shared/libs/price"
	grpc_ob "NFTM/shared/orderbook/grpc"
	"context"
	"fmt"
	"orderbook_service/lib/hub"
	"time"

	ob "github.com/i25959341/orderbook"
)

func (s *OrderbookServer) NewLimitOrder(ctx context.Context, args *grpc_ob.NewLimitOrderArgs) (*grpc_ob.NewLimitOrderResponse, error) {

	orderbook := hub.GetOderbookByPair(args.GetPair())

	quantity := price_helper.FromIntWithAppCoef(args.Quantity)
	price := price_helper.FromIntWithAppCoef(args.Price)

	done, partial, partialQuantityProcessed, err := orderbook.ProcessLimitOrder(FormatSide(args.Side), args.OrderId, quantity, price)
	transactionTime := time.Now()

	if err != nil {
		return nil, err
	}

	doneOrders := []*grpc_ob.Order{}

	for _, item := range done {
		doneOrders = append(doneOrders, ParseOrder(item))
	}

	response := &grpc_ob.NewLimitOrderResponse{
		Order:                    doneOrders,
		PartialQuantityProcessed: price_helper.ToIntWithAppCoef(&partialQuantityProcessed),
		TransactionTime:          transactionTime.Format(time.RFC3339),
	}

	if partial != nil {
		response.Partial = ParseOrder(partial)
	}

	return response, nil
}

func ParseOrder(input *ob.Order) *grpc_ob.Order {
	fmt.Printf("input: %v\n", input)

	quant := input.Quantity()
	price := input.Price()

	output := &grpc_ob.Order{
		Side:      grpc_ob.SideOrder(input.Side()),
		Id:        input.ID(),
		Timestamp: input.Time().Unix(),
		Quantity:  price_helper.ToIntWithAppCoef(&quant),
		Price:     price_helper.ToIntWithAppCoef(&price),
	}

	return output
}
