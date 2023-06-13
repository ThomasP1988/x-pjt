package config

import (
	"NFTM/shared/components/market"
	Pair "NFTM/shared/components/pair"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func SetEnv(env *Stage) *Stage {
	if env == nil {
		envS := Stage(os.Getenv("stage"))
		if envS != "" {
			return &envS

		} else {
			return &DEV

		}
	}
	return env
}

var Conf *Config

const infura string = ""

func GetConfig(env *Stage) *Config {

	if Conf != nil {
		return Conf
	}
	env = SetEnv(env)

	cryptoDAI := Pair.Pair{Base: "CRYPTO", Quote: "DAI"}

	Conf = &Config{
		AWSAccount: "",
		Region:     aws.String("eu-west-1"),
		Profile:    aws.String("nftquant"),
		Tables:     GetTables(env),
		Buckets:    GetBuckets(env),
		Markets: []market.MarketConfig{
			{
				Pair: cryptoDAI,
				Port: 40058,
				DNS:  cryptoDAI.StringLowercase(),
			},
		},
		User: ConfigCognito{
			UserPool:             "",
			UserPoolClient:       "",
			UserPoolClientSecret: "",
			UserPoolIdentity:     "",
			Domain:               "nftm-customer",
		},
		Admin: ConfigCognito{
			UserPool:         "",
			UserPoolClient:   "",
			UserPoolIdentity: "",
			Domain:           "nftm-admin",
		},
		Apis: map[Api]ApiConfig{
			MarketGRPC: {
				Port: 30001,
			},
			MarketGRPCHTTP: {
				Port: 30002,
			},
		},
		Infura: infura,
		Wallet: WalletConfig{
			Address: "",
		},
		SNS: SNSConfig{
			FetchItemTopic: string(*env) + "-FetchItem",
		},
		Blockchains: map[Blockchain]BlockchainConfig{
			Palm: {
				NodeURL: "https://palm-testnet.infura.io/v3/" + infura,
			},
			Ethereum: {
				NodeURL: "https://rinkeby.infura.io/v3/" + infura,
			},
		},
	}

	if *env == STAGING {
		updateConfigForStage(Conf)
	} else if *env == PROD {
		updateConfigForProd(Conf)
	}

	return Conf
}
