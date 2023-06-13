package order

import (
	entity "NFTM/shared/entities/order"
	dynamodb_helper "NFTM/shared/repositories/dynamodb"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type OrderListBySymbolArgs struct {
	Symbol *string
	IsOpen *int8
	Limit  *int32
	From   *string
}

func OrderListBySymbol(ctx context.Context, args OrderListBySymbolArgs) (*[]entity.Order, *string, error) {
	os := GetOrderService()
	newCond := expression.Key("symbol").Equal(expression.Value(*args.Symbol))
	if args.IsOpen != nil {
		newCond = newCond.And(expression.Key("isOpen").Equal(expression.Value(*args.IsOpen)))
	}
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
		IndexName:                 &os.SymbolIndex,
	}
	fmt.Printf("input: %v\n", input)
	if args.Limit != nil && *args.Limit > 0 {
		input.Limit = args.Limit
	}
	// fmt.Printf("args.From: %v\n", args.From)
	if args.From != nil && *args.From != "" {
		startKey, err := dynamodb_helper.Deserialize(*args.From, entity.Order{})
		if err != nil {
			return nil, nil, err
		}
		input.ExclusiveStartKey = *startKey
		fmt.Printf("input.ExclusiveStartKey: %v\n", input.ExclusiveStartKey)
	}

	output, err := os.Client.Query(ctx, &input)
	if err != nil {
		return nil, nil, err
	}

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
