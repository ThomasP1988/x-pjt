package collection

import (
	commonAws "NFTM/shared/common/aws"
	"NFTM/shared/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var service *CollectionService

type CollectionService struct {
	Client           *dynamodb.Client
	TableName        string
	SymbolIndex      string
	StatusIndex      string
	SubmittedAtIndex string
}

func GetCollectionService() *CollectionService {
	if service == nil {
		service = &CollectionService{
			Client:           commonAws.GetDynamoDBClient(),
			TableName:        config.Conf.Tables[config.COLLECTION].Name,
			SymbolIndex:      config.Conf.Tables[config.COLLECTION].SecondaryIndex[config.SymbolIndex],
			StatusIndex:      config.Conf.Tables[config.COLLECTION].SecondaryIndex[config.StatusIndex],
			SubmittedAtIndex: config.Conf.Tables[config.COLLECTION].SecondaryIndex[config.SubmittedAtIndex],
		}
	}

	return service
}
