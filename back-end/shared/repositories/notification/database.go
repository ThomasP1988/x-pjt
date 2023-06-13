package notification

import (
	"context"

	commonAws "NFTM/shared/common/aws"
	"NFTM/shared/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var service *NotificationService

type NotificationService struct {
	Client        *dynamodb.Client
	TableName     string
	UserDateIndex string
}

func GetNotificationService() *NotificationService {
	if service == nil {
		service = &NotificationService{
			Client:        commonAws.GetDynamoDBClient(),
			TableName:     config.Conf.Tables[config.NOTIFICATION].Name,
			UserDateIndex: config.Conf.Tables[config.NOTIFICATION].SecondaryIndex[config.UserDateIndex],
		}
	}

	return service
}

func Delete(context context.Context, userId string) error {
	ns := GetNotificationService()

	params := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"userId": &types.AttributeValueMemberS{
				Value: userId,
			},
		},
		TableName: &ns.TableName,
	}

	_, err := ns.Client.DeleteItem(context, params)

	if err != nil {
		return err
	}

	return nil
}
