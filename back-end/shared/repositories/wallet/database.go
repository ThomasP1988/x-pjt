package wallet

import (
	commonAws "NFTM/shared/common/aws"
	"NFTM/shared/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var service *WalletService

type WalletService struct {
	Client    *dynamodb.Client
	TableName string
}

func GetWalletService() *WalletService {
	if service == nil {
		service = &WalletService{
			Client:    commonAws.GetDynamoDBClient(),
			TableName: config.Conf.Tables[config.WALLET_ASSET].Name,
		}
	}

	return service
}
