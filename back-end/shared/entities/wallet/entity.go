package wallet

import "time"

const DefaultExp int32 = 8

type Wallet struct {
	UserID string
	Assets map[string]WalletAsset
}

type WalletAsset struct {
	UserID     string    `dynamodbav:"userId"`
	AssetID    string    `dynamodbav:"assetId"`
	LastUpdate time.Time `dynamodbav:"lastUpdate"`
	Own        int64     `dynamodbav:"own"`
	Available  int64     `dynamodbav:"available"`
}
