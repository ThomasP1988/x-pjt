package constants

type SideOrder int

const (
	Order_BUY  SideOrder = 0
	Order_SELL SideOrder = 1
)

type TypeOrder string

const (
	Order_MARKET TypeOrder = "market"
	Order_LIMIT  TypeOrder = "limit"
)

type StatusOrder string

const (
	Order_EMPTY            StatusOrder = "empty"
	Order_FILLED           StatusOrder = "filled"
	Order_CANCELLED        StatusOrder = "cancelled"
	Order_OPEN             StatusOrder = "open"
	Order_PARTIALLY_FILLED StatusOrder = "partially_filled"
)
