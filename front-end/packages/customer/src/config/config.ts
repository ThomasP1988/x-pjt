import StageConfig from "./stage";
import ProdConfig from "./prod";
import { Config } from "./types";
const env = process.env.REACT_APP_CUSTOM_ENV || 'dev';

const config: Config = {
	cognito: {
		REGION: "eu-west-1",
		USER_POOL_ID: "eu-west-1_uGAQWGowE",
		APP_CLIENT_ID: "1k0d1ct379h53ot23002hdfs5m",
		IDENTITY_POOL_ID: "eu-west-1:cab5a82e-40df-4833-a6c0-6bf783747589"
	},
	s3: {
		BUCKET: "dev-nftm-media",
		REGION: "eu-west-1"
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