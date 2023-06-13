package config

import "NFTM/shared/components/market"

type Stage string

var (
	DEV     Stage = "dev"
	STAGING Stage = "staging"
	PROD    Stage = "prod"
)

type Api int

const (
	MarketGRPCHTTP Api = iota
	MarketGRPC
)

type Config struct {
	AWSAccount  string
	Apis        map[Api]ApiConfig
	Region      *string
	Tables      TableRegistry
	Buckets     BucketRegistry
	Markets     []market.MarketConfig
	User        ConfigCognito
	Admin       ConfigCognito
	Profile     *string
	Infura      string
	Wallet      WalletConfig
	SNS         SNSConfig
	Blockchains map[Blockchain]BlockchainConfig
}

type ConfigCognito struct {
	UserPool             string
	UserPoolClient       string
	UserPoolIdentity     string
	UserPoolClientSecret string
	Domain               string
}

type ApiConfig struct {
	Port int
}

type WalletConfig struct {
	Address string
}

type SNSConfig struct {
	FetchItemTopic string
}
