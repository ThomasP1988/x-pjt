package nftitem

import (
	commonAws "NFTM/shared/common/aws"
	"NFTM/shared/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

var service *NFTItemService

type NFTItemService struct {
	Client        *dynamodb.Client
	SNSClient     *sns.Client
	TableName     string
	UserDateIndex string
}

func GetNFTItemService() *NFTItemService {
	if service == nil {
		service = &NFTItemService{
			Client:    commonAws.GetDynamoDBClient(),
			SNSClient: commonAws.GetSNSClient(),
			TableName: config.Conf.Tables[config.NFT_ITEM].Name,
		}
	}

	return service
}
