package orderbook

import (
	grpc_ob "NFTM/apis/grpc/expose/orderbook"
	"NFTM/apis/grpc/reducer"
	order_repo "NFTM/shared/repositories/order"
	auth_service "NFTM/shared/services/auth"
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func (s *OrderbookServer) CancelOrder(ctx context.Context, args *grpc_ob.CancelOrderArgs) (*grpc_ob.Order, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("error authentication")
	}
	fmt.Printf("args: %v\n", args)
	jwt := md.Get("jwt")[0]

	user, err := auth_service.GetUserService().Auth(jwt)
	if err != nil {
		return nil, errors.New("error authentication")
	}

	order, err := order_repo.Get(ctx, args.OrderId)
	if err != nil {
		return nil, err
	}

	err = OBServices.CancelOrder(ctx, order.Symbol, args.OrderId)

	if err != nil {
		return nil, err
	}

	if order.UserID != user.ID {
		return nil, errors.New("this order doesn't belong to connected user")
	}

	cancelledOrder, err := order_repo.Cancel(ctx, args.OrderId)

	if err != nil {
		return nil, err
	}

	return reducer.Order(cancelledOrder), nil
}
