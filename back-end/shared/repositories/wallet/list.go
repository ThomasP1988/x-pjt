package wallet

import (
	entity "NFTM/shared/entities/wallet"
	dynamodb_helper "NFTM/shared/repositories/dynamodb"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ListArgs struct {
	UserID string
	Limit  *int32
	From   *string
}

func List(ctx context.Context, args ListArgs) (*[]entity.WalletAsset, *string, error) {
	ws := GetWalletService()

	newCond := expression.Key("userId").Equal(expression.Value(args.UserID))
	expr, err := expression.NewBuilder().WithKeyCondition(newCond).Build()

	if err != nil {
		return nil, nil, err
	}

	orders := &[]entity.WalletAsset{}

	input := dynamodb.QueryInput{
		TableName:                 &ws.TableName,
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ScanIndexForward:          aws.Bool(false),
	}

	if args.Limit != nil && *args.Limit > 0 {
		input.Limit = args.Limit
	}
	if args.From != nil && *args.From != "" {
		startKey, err := dynamodb_helper.Deserialize(*args.From, entity.WalletAsset{})
		if err != nil {
			return nil, nil, err
		}
		input.ExclusiveStartKey = *startKey
	}

	output, err := ws.Client.Query(ctx, &input)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("output: %v\n", output)
	err = attributevalue.UnmarshalListOfMaps(output.Items, orders)
	if err != nil {
		return nil, nil, err
	}

	next, err := dynamodb_helper.Serialize(output.LastEvaluatedKey, &entity.WalletAsset{})
	if err != nil {
		return nil, nil, err
	}

	return orders, next, nil
}
