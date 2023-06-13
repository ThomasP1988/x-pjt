package reducer

import (
	grpc_ob "NFTM/apis/grpc/expose/orderbook"
	"NFTM/shared/constants"
)

func TypeOrder(input *constants.TypeOrder) *grpc_ob.TypeOrder {
	if *input == constants.Order_LIMIT {
		return grpc_ob.TypeOrder_LIMIT.Enum()
	} else {
		return grpc_ob.TypeOrder_MARKET.Enum()
	}
}
