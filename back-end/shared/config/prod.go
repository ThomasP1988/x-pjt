package config

func updateConfigForProd(conf *Config) {
	conf.Blockchains = map[Blockchain]BlockchainConfig{
		Palm: {
			NodeURL: "https://palm-mainnet.infura.io/v3/" + infura,
		},
		Ethereum: {
			NodeURL: "https://mainnet.infura.io/v3/" + infura,
		},
	}
}
