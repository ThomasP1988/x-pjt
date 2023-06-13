package main

import (
	"NFTM/shared/config"
	"time"

	ob_entity "NFTM/shared/orderbook/entities"
	"math/rand"

	"github.com/google/uuid"
	ob "github.com/i25959341/orderbook"
	"github.com/shopspring/decimal"
)

func main() {
	// Start()
	config.GetConfig(nil)
	StartGRPC()
	// GenerateTestData()
}

func GenerateTestData() {
	for {
		time.Sleep(time.Second * 2)
		uidBuy, _ := uuid.NewUUID()
		ob_entity.Orderbook.ProcessLimitOrder(ob.Buy, uidBuy.String(), decimal.New(int64(rand.Intn(10000)), 0), decimal.New(int64(rand.Intn(10000)), 0))
		// done, partial, quantityProce
		// fmt.Printf("done: %v\n", done)
		// fmt.Printf("partial: %v\n", partial)
		// fmt.Printf("quantityProcessed: %v\n", quantityProcessed)
		// fmt.Printf("err: %v\n", err)
		uidSell, _ := uuid.NewUUID()

		ob_entity.Orderbook.ProcessLimitOrder(ob.Sell, uidSell.String(), decimal.New(int64(rand.Intn(10000)), 0), decimal.New(int64(rand.Intn(10000)), 0))
		// done, partial, quantityProcessed, err = Orderbook.ProcessLimitOrder(ob.Sell, uidSell.String(), decimal.New(int64(rand.Intn(10000)), 0), decimal.New(int64(rand.Intn(10000)), 0))
		// fmt.Printf("done: %v\n", done)
		// fmt.Printf("partial: %v\n", partial)
		// fmt.Printf("quantityProcessed: %v\n", quantityProcessed)
		// fmt.Printf("err: %v\n", err)
	}
}
