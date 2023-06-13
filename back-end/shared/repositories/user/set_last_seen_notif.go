package user

import (
	entity "NFTM/shared/entities/user"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func SetLastSeenNotification(ctx context.Context, userID string, lastSeenTime time.Time) (*entity.User, error) {
	us := GetUserService()
	output, err := us.Client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: &us.TableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: userID},
		},
		UpdateExpression: aws.String("set lastSeenNotification = :lastSeenNotification"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":lastSeenNotification": &types.AttributeValueMemberS{Value: lastSeenTime.Format(time.RFC3339)},
		},
		ReturnValues: types.ReturnValueAllNew,
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	updatedUser := &entity.User{}
	attributevalue.UnmarshalMap(output.Attributes, updatedUser)
	return updatedUser, nil
}
