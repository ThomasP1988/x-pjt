package order

import (
	commonAws "NFTM/shared/common/aws"
	"NFTM/shared/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var service *OrderService

type OrderService struct {
	Client              *dynamodb.Client
	TableName           string
	UserIndex           string
	SymbolIndex         string
	UserSymbolOpenIndex string
}

func GetOrderService() *OrderService {
	if service == nil {

		service = &OrderService{
			Client:              commonAws.GetDynamoDBClient(),
			TableName:           config.Conf.Tables[config.ORDER].Name,
			UserIndex:           config.Conf.Tables[config.ORDER].SecondaryIndex[config.UserIndex],
			SymbolIndex:         config.Conf.Tables[config.ORDER].SecondaryIndex[config.SymbolIndex],
			UserSymbolOpenIndex: config.Conf.Tables[config.ORDER].SecondaryIndex[config.UserIdSymbolIsOpenIndex],
		}
	}

	return service
}
