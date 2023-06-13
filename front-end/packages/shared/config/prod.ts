import { Blockchains, Config } from "./types";

const prod: Partial<Config> = {
	blockchain: {
		[Blockchains.ETHEREUM]: {
			name: "Ethereum",
			chainId: 1,
			network: "",
			explorer: "",
			svg: "/img/ethereum-eth-logo.svg"
		},
		[Blockchains.POLYGON]: {
			name: "Polygon Mainnet",
			chainId: 137,
			network: "https://polygon-rpc.com",
			explorer: "https://polygonscan.com",
			currency: {
				name: "MATIC",
				symbol: "MATIC",
				decimals: 18
			},
			svg: ""
		},
		[Blockchains.PALM]: {
			name: "Palm Testnet",
			chainId: 0,
			network: "",
			explorer: "",
			currency: {
				name: "PALM",
				symbol: "PALM",
				decimals: 18
			},
			svg: ""
		},
	}
}

export default prod;