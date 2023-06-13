package order

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Delete(ctx context.Context, orderId string) error {
	os := GetOrderService()
	_, err := os.Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"orderId": &types.AttributeValueMemberS{
				Value: orderId,
			},
		},
		TableName: &os.TableName,
	})

	return err
}
