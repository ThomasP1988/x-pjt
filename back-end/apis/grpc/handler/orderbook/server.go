package orderbook

import (
	grpc_ob "NFTM/apis/grpc/expose/orderbook"
	"NFTM/shared/services/orderbooks"
)

var OBServices *orderbooks.OrderbooksService

type OrderbookServer struct {
	grpc_ob.OrderbooksServer
}

func NewOrderbookServer() *OrderbookServer {
	OBServices = orderbooks.NewOrderbooksService()
	return &OrderbookServer{}
}
