package server

import (
	grpc_ob "NFTM/shared/orderbook/grpc"

	ob "github.com/i25959341/orderbook"
)

func FormatSide(side grpc_ob.SideOrder) ob.Side {
	if side == grpc_ob.SideOrder_BUY { // the creator of the oderbook library made a mistake between 0 and 1 buy/sell
		return ob.Buy
	} else {
		return ob.Sell
	}
}
