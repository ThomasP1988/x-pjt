package server

import (
	grpc_ob "NFTM/shared/orderbook/grpc"
	"context"
	"orderbook_service/lib/hub"
)

func (s *OrderbookServer) CancelOrder(ctx context.Context, args *grpc_ob.CancelOrderArgs) (*grpc_ob.CancelOrderResult, error) {

	orderbook := hub.GetOderbookByPair(args.GetPair())

	orderbook.CancelOrder(args.GetOrderId())

	return &grpc_ob.CancelOrderResult{
		Success: true,
	}, nil
}
