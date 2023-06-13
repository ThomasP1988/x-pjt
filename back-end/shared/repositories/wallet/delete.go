package wallet

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Delete(ctx context.Context, userID string, assetID string) error {
	ws := GetWalletService()
	_, err := ws.Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"userId": &types.AttributeValueMemberS{
				Value: userID,
			},
			"assetId": &types.AttributeValueMemberS{
				Value: assetID,
			},
		},
		TableName: &ws.TableName,
	})

	return err
}
