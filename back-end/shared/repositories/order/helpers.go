package order

import (
	"strconv"
)

func FormatUserIDSymbolIsOpen(userId string, symbol string, isOpen int8) string {
	return userId + symbol + strconv.Itoa(int(isOpen))
}
