package order

import (
	entity "NFTM/shared/entities/order"
	dynamodb_helper "NFTM/shared/repositories/dynamodb"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type OrderListArgs struct {
	UserID *string
	Symbol *string
	IsOpen int8
	Limit  *int32
	From   *string
}

func OrderList(ctx context.Context, args OrderListArgs) (*[]entity.Order, *string, error) {
	os := GetOrderService()
	newCond := expression.Key("userIdSymbolIsOpen").Equal(expression.Value(FormatUserIDSymbolIsOpen(*args.UserID, *args.Symbol, args.IsOpen)))
	expr, err := expression.NewBuilder().WithKeyCondition(newCond).Build()

	if err != nil {
		return nil, nil, err
	}

	orders := &[]entity.Order{}

	input := dynamodb.QueryInput{
		TableName:                 &os.TableName,
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		IndexName:                 &os.UserSymbolOpenIndex,
		ScanIndexForward:          aws.Bool(false),
	}

	if args.Limit != nil && *args.Limit > 0 {
		input.Limit = args.Limit
	}
	if args.From != nil && *args.From != "" {
		startKey, err := dynamodb_helper.Deserialize(*args.From, entity.Order{})
		if err != nil {
			return nil, nil, err
		}
		input.ExclusiveStartKey = *startKey
	}

	output, err := os.Client.Query(ctx, &input)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("output: %v\n", output)
	err = attributevalue.UnmarshalListOfMaps(output.Items, orders)
	if err != nil {
		return nil, nil, err
	}

	next, err := dynamodb_helper.Serialize(output.LastEvaluatedKey, &entity.Order{})
	if err != nil {
		return nil, nil, err
	}

	return orders, next, nil
}
