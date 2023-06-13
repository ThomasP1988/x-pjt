package server

import (
	grpc_ob "NFTM/shared/orderbook/grpc"
	"orderbook_service/lib/hub"
	"sort"
	"time"

	ob "github.com/i25959341/orderbook"
)

func (s *OrderbookServer) SubscribeL2(args *grpc_ob.SubscribeArgs, send grpc_ob.Orderbook_SubscribeL2Server) error {

	orderbook := hub.GetOderbookByPair(args.GetPair())

	for {
		asks, bids := orderbook.Depth()
		// fmt.Printf("asks: %v\n", len(asks))
		// fmt.Printf("bids: %v\n", len(bids))

		marshalledAsks, hasAskUpdate := HandlePriceLevelAsks(asks, SizeL2)
		marshalledBids, hasBidUpdate := HandlePriceLevelBids(bids, SizeL2)

		if hasAskUpdate || hasBidUpdate {
			send.Send(&grpc_ob.Prices{
				Asks: marshalledAsks[:],
				Bids: marshalledBids[:],
			})
		}

		time.Sleep(time.Millisecond * 300)
	}
}

func HandlePriceLevelAsks(input []*ob.PriceLevel, depth int) ([SizeL2]*grpc_ob.PriceLevel, bool) {
	output := [SizeL2]*grpc_ob.PriceLevel{}

	if depth > len(input) {
		depth = len(input)
	}

	sort.Slice(input, func(i, j int) bool {
		return input[i].Price.LessThan(input[j].Price)
	})

	// ASKS input ordered forward, $depth first results

	for i := 0; i < depth; i++ {
		if input[i] != nil {
			output[i] = &grpc_ob.PriceLevel{
				Price:    input[i].Price.String(),
				Quantity: input[i].Quantity.String(),
			}
		} else {
			output[i] = nil
		}
	}
	return output, true
}

func HandlePriceLevelBids(input []*ob.PriceLevel, depth int) ([SizeL2]*grpc_ob.PriceLevel, bool) {
	output := [SizeL2]*grpc_ob.PriceLevel{}

	if depth > len(input) {
		depth = len(input)
	}

	sort.Slice(input, func(i, j int) bool {
		return input[i].Price.GreaterThan(input[j].Price)
	})

	// BIDS input ordered backward, $depth first results
	for i := 0; i < depth; i++ {
		if input[i] != nil {
			output[i] = &grpc_ob.PriceLevel{
				Price:    input[i].Price.String(),
				Quantity: input[i].Quantity.String(),
			}
		} else {
			output[i] = nil
		}
	}

	return output, true
}
