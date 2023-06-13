package reducer

import (
	grpc_ob "NFTM/apis/grpc/expose/orderbook"
	order_entity "NFTM/shared/entities/order"
	price_helper "NFTM/shared/libs/price"
)

func Order(input *order_entity.Order) *grpc_ob.Order {
	return &grpc_ob.Order{
		Id:               input.ID,
		Status:           string(input.Status),
		Symbol:           input.Symbol,
		Side:             grpc_ob.SideOrder(input.Side),
		Type:             *TypeOrder(&input.Type),
		CreatedAt:        input.CreatedAt.UnixMilli(),
		Price:            price_helper.FromIntToString(input.Price),
		Quantity:         price_helper.FromIntToString(input.Quantity),
		OriginalQuantity: price_helper.FromIntToString(input.OriginalQuantity),
		FilledQuantity:   price_helper.FromIntToString(input.FilledQuantity),
	}
}
