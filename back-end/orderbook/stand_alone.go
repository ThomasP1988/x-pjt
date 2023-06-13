package main

import (
	ob_helpers "NFTM/shared/orderbook/helpers"
	"fmt"
	"orderbook_service/lib/hub"
	"os"

	ob "github.com/i25959341/orderbook"
	"google.golang.org/grpc"
)

func StartStandAlone(apiserver *grpc.Server) {
	hub.IsStandAlone = true

	symbol := os.Getenv("MARKET_SYMBOL")
	fmt.Printf("symbol: %v\n", symbol)
	// symbol = "CRYPTO-DAI"
	hub.OrderBookStandAlone = ob.NewOrderBook()

	ob_helpers.FillOrderbook(hub.OrderBookStandAlone, symbol)

}
