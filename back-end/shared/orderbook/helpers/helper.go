package ob_helpers

import (
	"NFTM/shared/constants"

	ob "github.com/i25959341/orderbook"
)

func FormatSide(side constants.SideOrder) ob.Side {
	if side == constants.Order_BUY { // the creator of the oderbook library made a mistake between 0 and 1 buy/sell
		return ob.Buy
	} else {
		return ob.Sell
	}
}
