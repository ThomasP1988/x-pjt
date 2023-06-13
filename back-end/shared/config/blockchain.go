package config

type Blockchain string
type BlockchainConfig struct {
	NodeURL string
}

const (
	Ethereum Blockchain = "ethereum"
	Palm     Blockchain = "palm"
	Polygon  Blockchain = "polygon"
)
