package config

import "os"

type Table int
type SecondaryIndex int

type SecondaryIndexes = map[SecondaryIndex]string

type TableConfig struct {
	Name           string
	SecondaryIndex SecondaryIndexes
}
type TableRegistry map[Table]TableConfig

const (
	USER Table = iota
	GRAPHQL_SUB_CONNECTION
	GRAPHQL_SUB_SUBSCRIPTION
	GRAPHQL_SUB_EVENT
	WALLET_ASSET
	ORDER
	PAYMENT_EVENT
	NOTIFICATION
	COLLECTION
	NFT_ITEM
	TOKEN
	BLOCKCHAIN_WALLET
)

const (
	EmailIndex SecondaryIndex = iota
	OperationIndex
	UserIndex
	SymbolIndex
	UserIdSymbolIsOpenIndex
	UserDateIndex
	StatusIndex
	SubmittedAtIndex
)

func GetTables(env *Stage) TableRegistry {

	SetEnv(env)

	if env == nil {
		envS := Stage(os.Getenv("env"))
		env = &envS
	}

	return TableRegistry{
		USER: TableConfig{
			Name: string(*env) + "-user",
			SecondaryIndex: SecondaryIndexes{
				EmailIndex: "email-index",
			},
		},
		WALLET_ASSET: TableConfig{
			Name: string(*env) + "-wallet-asset",
		},
		ORDER: TableConfig{
			Name: string(*env) + "-order",
			SecondaryIndex: SecondaryIndexes{
				UserIndex:               "user-index",
				SymbolIndex:             "symbol-index",
				UserIdSymbolIsOpenIndex: "user-symbol-open-index",
			},
		},
		PAYMENT_EVENT: TableConfig{
			Name: string(*env) + "-payment-event",
			SecondaryIndex: SecondaryIndexes{
				UserIndex: "user-index",
			},
		},
		GRAPHQL_SUB_CONNECTION: TableConfig{
			Name: string(*env) + "-ws-connection",
		},
		GRAPHQL_SUB_SUBSCRIPTION: TableConfig{
			Name: string(*env) + "-ws-subscription",
			SecondaryIndex: SecondaryIndexes{
				OperationIndex: "operation-index",
			},
		},
		GRAPHQL_SUB_EVENT: TableConfig{
			Name: string(*env) + "-ws-event",
		},
		COLLECTION: TableConfig{
			Name: string(*env) + "-collection",
			SecondaryIndex: SecondaryIndexes{
				SymbolIndex:      "symbol-index",
				StatusIndex:      "status-index",
				SubmittedAtIndex: "submittedAt-index",
			},
		},
		NFT_ITEM: TableConfig{
			Name: string(*env) + "-nft-item",
		},
		NOTIFICATION: TableConfig{
			Name: string(*env) + "-notification",
			SecondaryIndex: SecondaryIndexes{
				UserDateIndex: "user-date-index",
			},
		},
		TOKEN: TableConfig{
			Name: string(*env) + "-token",
		},
		BLOCKCHAIN_WALLET: TableConfig{
			Name: string(*env) + "-blockchain-wallet",
		},
	}
}
