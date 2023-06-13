package orderbooks

func (os *OrderbooksService) HasPair(pair string) bool {
	_, ok := os.OrderbookClients[pair]
	return ok
}
