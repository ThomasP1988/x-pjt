package market

import (
	pair "NFTM/shared/components/pair"
)

type MarketConfig struct {
	Pair pair.Pair
	Port int
	DNS  string
}
