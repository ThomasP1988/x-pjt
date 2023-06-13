package orderbooks

import (
	pair_component "NFTM/shared/components/pair"
	"NFTM/shared/config"
	ob_client "NFTM/shared/orderbook/client"
	"NFTM/shared/repositories/log"
	"fmt"
	"strconv"
	"time"
)

func (os *OrderbooksService) Start(pair string) (*pair_component.Pair, error) {

	if os.HasPair(pair) {
		fmt.Printf("\n pair: %v\n", pair)
		fmt.Printf("\n os.OrderbookClients[pair]: %v\n", *os.OrderbookClients[pair])
		return os.OrderbookClients[pair].Pair, nil
	}

	conf := config.GetConfig(nil)
	market := conf.Markets[0] //TODO: the markets should be load by the DB, not the config
	address := market.DNS + ":" + strconv.Itoa(market.Port)

	pairStruct := conf.Markets[0].Pair

	for {
		newClient, err := ob_client.NewOrderbookStandAloneClient(address, pairStruct)

		if err != nil {
			fmt.Printf("can not connect orderbook: %v\n", err)
			log.Error("can't start new orderbook client, orderbooks/start", err)
		} else {
			os.OrderbookClients[pair] = newClient
			break
		}

		time.Sleep(time.Second)
	}

	return &pairStruct, nil
}
