package server

import (
	grpc_ob "NFTM/shared/orderbook/grpc"
)

// server is used to implement helloworld.GreeterServer.
type OrderbookServer struct {
	grpc_ob.OrderbookServer
}

const SizeL2 int = 10
