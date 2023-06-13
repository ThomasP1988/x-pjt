package orderbooks

import (
	ob_client "NFTM/shared/orderbook/client"
)

type OrderbooksService struct {
	OrderbookClients map[string]*ob_client.OrderbookStandAloneClient
}

func NewOrderbooksService() *OrderbooksService {
	return &OrderbooksService{
		OrderbookClients: map[string]*ob_client.OrderbookStandAloneClient{},
	}

}
