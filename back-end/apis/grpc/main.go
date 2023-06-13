package main

import (
	"NFTM/apis/grpc/handler/orderbook"
	"NFTM/shared/config"
	"fmt"
	"os"
)

func main() {
	config.GetConfig(nil)
	orderbook.StartBrokerL2(&map[string]*orderbook.Broker{})
	fmt.Println(os.Environ())

	go StartGRPC()
	StartGRPCHTTP()
}
