package token

import (
	commonAws "NFTM/shared/common/aws"
	"NFTM/shared/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var service *TokenService

type TokenService struct {
	Client    *dynamodb.Client
	TableName string
}

func GetTokenService() *TokenService {
	if service == nil {

		service = &TokenService{
			Client:    commonAws.GetDynamoDBClient(),
			TableName: config.Conf.Tables[config.TOKEN].Name,
		}
	}

	return service
}
