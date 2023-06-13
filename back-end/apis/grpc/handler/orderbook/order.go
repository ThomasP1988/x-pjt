package orderbook

import (
	grpc_ob "NFTM/apis/grpc/expose/orderbook"
	"context"
	"errors"
	"fmt"

	"NFTM/shared/constants"
	price_helper "NFTM/shared/libs/price"
	client_ob "NFTM/shared/orderbook/grpc"

	auth_service "NFTM/shared/services/auth"
	"NFTM/shared/services/orderbooks"

	"google.golang.org/grpc/metadata"
)

func (s *OrderbookServer) ProcessOrder(ctx context.Context, args *grpc_ob.OrderArgs) (*grpc_ob.Order, error) {

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

	side := client_ob.SideOrder(args.GetSide())
	pair, err := OBServices.Start(args.GetSymbol())

	if err != nil {
		return nil, err
	}

	// tech debt: grpc web sends null instead of 0, so if null let's set to 0

	fmt.Printf("args side: %v\n", args.Side)
	fmt.Printf("side: %v\n", side)
	fmt.Printf("type: %v\n", args.Type)

	var orderResult *orderbooks.OrderResult
	// fmt.Printf("\"test\": %v\n", "test")
	if args.Type == grpc_ob.TypeOrder_LIMIT {
		orderResult, err = OBServices.ProcessLimitOrder(ctx, user.ID, pair, side, args.GetQuantity(), args.GetPrice())
	}

	if args.Type == grpc_ob.TypeOrder_MARKET {
		orderResult, err = OBServices.ProcessMarketOrder(ctx, user.ID, pair, side, args.GetQuantity())
	}

	if err != nil {
		return nil, err
	}

	var typeOrder grpc_ob.TypeOrder

	if orderResult.Type == constants.Order_LIMIT {
		typeOrder = grpc_ob.TypeOrder_LIMIT
	} else {
		typeOrder = grpc_ob.TypeOrder_MARKET
	}

	result := grpc_ob.Order{
		Id:               orderResult.OrderId,
		CreatedAt:        orderResult.TransactionTime.Unix(),
		Price:            price_helper.FromIntToString(*orderResult.Price),
		OriginalQuantity: price_helper.FromIntToString(*orderResult.OriginalQuantity),
		Paid:             price_helper.FromIntToString(*orderResult.CumulativeQuoteQuantity),
		Symbol:           orderResult.Symbol,
		AveragePrice:     price_helper.FromIntToString(*orderResult.Price),
		Quantity:         price_helper.FromIntToString(*orderResult.ExecutedQuantity),
		Status:           string(orderResult.Status),
		Type:             typeOrder,
		Side:             grpc_ob.SideOrder(orderResult.Side),
	}
	return &result, nil

}
