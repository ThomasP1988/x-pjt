import StageConfig from "./stage";
import ProdConfig from "./prod";
import { Config } from "./types";
const env = process.env.REACT_APP_CUSTOM_ENV || 'dev';

const config: Config = {
	cognito: {
		REGION: "eu-west-1",
		USER_POOL_ID: "eu-west-1_0O9TKYnck",
		APP_CLIENT_ID: "74kua5ogtnckuotq563j3fupqa",
		IDENTITY_POOL_ID: "eu-west-1:8cb1c96c-a59c-43bd-a1e5-8a212e0e3455"
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