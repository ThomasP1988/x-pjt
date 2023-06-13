import StageConfig from "./stage";
import ProdConfig from "./prod";
import { Blockchains, Config } from "./types";
const env = process.env.REACT_APP_CUSTOM_ENV || 'dev';

const config: Config = {
	// s3: {
	//   REGION: "eu-west-1",
	//   BUCKET: "leviathan-attachment-dev",
	// },
	apiGateway: {
		REGION: 'us-east-1',
		HTTP: 'https://e8fl3xf0f6.execute-api.eu-west-1.amazonaws.com/graphql',
		WS: 'wss://7twsv473bj.execute-api.eu-west-1.amazonaws.com/dev'
	},
	appsync: {
		REGION: 'eu-west-1',
		HTTP: "https://f65dq4p2vjapvpqtuz626to2ua.appsync-api.eu-west-1.amazonaws.com/graphql",
		WS: "wss://f65dq4p2vjapvpqtuz626to2ua.appsync-realtime-api.eu-west-1.amazonaws.com/graphql",
	},
	blockchain: {
		[Blockchains.ETHEREUM]: {
			name: "Ethereum testnet",
			chainId: 3,
			network: "https://rinkeby-light.eth.linkpool.io",
			explorer: "https://rinkeby.etherscan.io",
			currency: {
				name: "RIN",
				symbol: "RIN",
				decimals: 18
			},
			svg: "/img/ethereum-eth-logo.svg"
		},
		[Blockchains.POLYGON]: {
			name: "Polygon Testnet Mumbai",
			chainId: 80001,
			network: "https://rpc-mumbai.matic.today",
			explorer: "https://mumbai.polygonscan.com",
			currency: {
				name: "MATIC",
				symbol: "MATIC",
				decimals: 18
			},
			svg: "/img/polygon-matic-logo.svg"
		},
		[Blockchains.PALM]: {
			name: "Palm Testnet",
			chainId: 11297108099,
			network: "https://palm-testnet.infura.io/v3/",
			explorer: "https://explorer.palm-uat.xyz",
			currency: {
				name: "PALM",
				symbol: "PALM",
				decimals: 18
			},
			svg: "/img/immutable-x-imx-logo.svg"
		},
	}
};

switch (env) {
	case 'stage':
		Object.assign(config, StageConfig);
		break;
	case 'prod':
		Object.assign(config, ProdConfig);
		break;
	case 'dev':
	default:
		Object.assign(config, {});
}


export default config;