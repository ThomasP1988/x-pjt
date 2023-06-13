package user

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Delete(context context.Context, userID string) error {
	us := GetUserService()
	_, err := us.Client.DeleteItem(context, &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: userID,
			},
		},
		TableName: &us.TableName,
	})

	return err

}
