package token

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Use(ctx context.Context, tokenID string, token interface{}) error {
	ts := GetTokenService()
	output, err := ts.Client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &ts.TableName,
		Key: map[string]types.AttributeValue{
			"tokenId": &types.AttributeValueMemberS{
				Value: tokenID,
			},
		},
	})

	if err != nil {
		return err
	}

	err = attributevalue.UnmarshalMap(output.Item, token)

	if err != nil {
		return err
	}

	err = Delete(ctx, tokenID)

	if err != nil {
		fmt.Printf("err deleting token: %v\n", err)
	}

	return nil

}
