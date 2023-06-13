package orderbooks

import (
	"NFTM/shared/constants"
	orderbook "NFTM/shared/orderbook/grpc"
	"time"
)

type OrderResult struct {
	Symbol                  string
	OrderId                 string
	TransactionTime         *time.Time
	Price                   *int64
	OriginalQuantity        *int64
	ExecutedQuantity        *int64
	CumulativeQuoteQuantity *int64
	Status                  constants.StatusOrder
	Type                    constants.TypeOrder
	Side                    orderbook.SideOrder
}
