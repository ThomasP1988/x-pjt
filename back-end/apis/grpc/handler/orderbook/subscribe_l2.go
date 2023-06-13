package orderbook

import (
	grpc_ob "NFTM/apis/grpc/expose/orderbook"
	orderbook "NFTM/shared/orderbook/grpc"
	"fmt"
)

var brokersL2 *map[string]*Broker
var publishCallback = func(pair string, price *orderbook.Prices) {
	(*brokersL2)[pair].Publish(price)
}

func StartBrokerL2(brokersL2Arg *map[string]*Broker) {
	brokersL2 = brokersL2Arg
}

func (s *OrderbookServer) SubscribeL2(args *grpc_ob.SubscribeL2Args, send grpc_ob.Orderbooks_SubscribeL2Server) error {
	fmt.Printf("\"Subscribe\": %v\n", args.GetSymbol())
	_, err := OBServices.Start(args.GetSymbol())

	if err != nil {
		fmt.Printf("error starting orderbook (GRPC API): %v\n", err)
		return err
	}

	if _, ok := (*brokersL2)[args.GetSymbol()]; !ok {
		fmt.Printf("\"1\": %v\n", "1")
		(*brokersL2)[args.GetSymbol()] = NewBroker(args.GetSymbol())
		go (*brokersL2)[args.GetSymbol()].Start()
	}

	OBServices.StartL2(args.GetSymbol(), &publishCallback)

	fmt.Printf("\"3\": %v\n", "3")
	c := (*brokersL2)[args.GetSymbol()].Subscribe()
	defer (*brokersL2)[args.GetSymbol()].Unsubscribe(c)
	fmt.Printf("\"4\": %v\n", "4")
	for {
		select {
		case <-send.Context().Done():
			return nil

		case price := <-c:
			// fmt.Printf("price.Asks: %v\n", price.Asks)
			// fmt.Printf("price.Bids: %v\n", price.Bids)
			err = send.SendMsg(price)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return nil
			}

		}
	}

}
