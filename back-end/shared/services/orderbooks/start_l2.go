package orderbooks

import (
	orderbook "NFTM/shared/orderbook/grpc"
	"fmt"
	"io"
)

var chanL2 map[string]chan interface{} = map[string]chan interface{}{}

func (os *OrderbooksService) StopL2(pair string) {
	chanL2[pair] <- true
	delete(chanL2, pair)

	fmt.Printf("stoping l2: %v\n", chanL2)
}

func (os *OrderbooksService) StartL2(pair string, onMessage *func(pair string, price *orderbook.Prices)) error {

	if _, ok := chanL2[pair]; ok {
		fmt.Printf("\"L2 Subscription\": %v\n", "already exists")
		return nil
	}

	// this function should be triggered only once (except if every has unsubscribed)
	// we start the service to get the prices from the orderbook (standalone or pool)
	// and publish the prices to our broker which publishes it to each connection
	go func() {
		stream, err := os.OrderbookClients[pair].StartL2()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		chanL2[pair] = make(chan interface{})
		for {
			select {
			case <-chanL2[pair]:
				return
			default:
				price, err := (*stream).Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Printf("err: %v\n", err)
				}

				(*onMessage)(pair, price)
			}

		}
	}()

	println("la")
	return nil
}
