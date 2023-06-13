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

func (s *OrderbookServer) OrderList(ctx context.Context, args *grpc_ob.OrderListArgs) (*grpc_ob.OrderListResult, error) {
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
	fmt.Printf("2args.From: %v\n", args.From)

	var isOpenInt int8

	if args.IsOpen {
		isOpenInt = 1
	} else {
		isOpenInt = 0
	}

	symbol := args.GetSymbol()
	orders, next, err := order_repo.OrderList(ctx, order_repo.OrderListArgs{
		UserID: &user.ID,
		Symbol: &symbol,
		IsOpen: isOpenInt,
		Limit:  &args.Limit,
		From:   &args.From,
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	var ordersFormated []*grpc_ob.Order

	for _, order := range *orders {
		ordersFormated = append(ordersFormated, reducer.Order(&order))
	}
	fmt.Printf("next: %v\n", *next)
	result := grpc_ob.OrderListResult{
		Next:   *next,
		Orders: ordersFormated,
	}

	return &result, nil
}
