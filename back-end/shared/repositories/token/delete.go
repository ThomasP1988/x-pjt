package token

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Delete(ctx context.Context, tokenId string) error {
	ts := GetTokenService()
	_, err := ts.Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"tokenId": &types.AttributeValueMemberS{
				Value: tokenId,
			},
		},
		TableName: &ts.TableName,
	})

	return err
}
