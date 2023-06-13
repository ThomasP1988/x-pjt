package user

import (
	commonAws "NFTM/shared/common/aws"
	"NFTM/shared/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var service *UserService

type UserService struct {
	Client     *dynamodb.Client
	TableName  string
	UserIndex  string
	EmailIndex string
}

func GetUserService() *UserService {
	if service == nil {

		service = &UserService{
			Client:     commonAws.GetDynamoDBClient(),
			TableName:  config.Conf.Tables[config.USER].Name,
			EmailIndex: config.Conf.Tables[config.USER].SecondaryIndex[config.EmailIndex],
		}
	}

	return service
}
