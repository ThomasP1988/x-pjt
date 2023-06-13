package ob_entity

import (
	"fmt"

	"github.com/google/uuid"
	ob "github.com/i25959341/orderbook"
	"github.com/shopspring/decimal"
)

var Orderbook *ob.OrderBook

func Start() {
	Orderbook = ob.NewOrderBook()
}

func Stop() {
	Orderbook = nil
}

func Buy() {
	// TODO: 1. handle limit order 2. send to influx db if trade
	uidBuy, _ := uuid.NewUUID()
	done, partial, quantityProcessed, err := Orderbook.ProcessLimitOrder(ob.Buy, uidBuy.String(), decimal.New(7, 0), decimal.New(120, 0))
	fmt.Printf("done: %v\n", done)
	fmt.Printf("partial: %v\n", partial)
	fmt.Printf("quantityProcessed: %v\n", quantityProcessed)
	fmt.Printf("err: %v\n", err)
}

func Sell() {
	// TODO: 1. handle limit order 2. send to influx db if trade
	uidSell, _ := uuid.NewUUID()

	done, partial, quantityProcessed, err := Orderbook.ProcessLimitOrder(ob.Sell, uidSell.String(), decimal.New(55, 0), decimal.New(100, 0))
	fmt.Printf("done: %v\n", done)
	fmt.Printf("partial: %v\n", partial)
	fmt.Printf("quantityProcessed: %v\n", quantityProcessed)
	fmt.Printf("err: %v\n", err)
}
