package hub

import (
	"fmt"
	"sync"

	ob "github.com/i25959341/orderbook"
)

var IsStandAlone bool = false
var OrderBookStandAlone *ob.OrderBook

func GetOderbookByPair(pair string) *ob.OrderBook {

	if IsStandAlone {
		return OrderBookStandAlone
	} else {
		// get in pool

		// if no, fetch orders in db and put a new orderbook in pool

		// return orderbook

	}

	return nil
}

func SetPool() {
	test := sync.Pool{
		New: func() interface{} {
			return OrderBookStandAlone
		},
	}

	fmt.Printf("test: %v\n", test)
}
