package candlestick

import "time"

type TypeRange string

// _start:2022-01-07 09:37:52.6928168 +0000 UTC,_stop:2022-01-07 10:37:52.6928168 +0000 UTC,_measurement:trades,pair:CRYPTO-DAI,table:3,_time:2022-01-07 10:37:42.656326 +0000 UTC
type Candlestick struct {
	ID     string    `json:"tradeId"`
	Start  time.Time `json:"_start"`
	Stop   time.Time `json:"_stop"`
	Symbol string    `json:"pair"`
	Side   string    `json:"side"`
	Price  string    `json:"price"`
	Amount string    `json:"amount"`
}
