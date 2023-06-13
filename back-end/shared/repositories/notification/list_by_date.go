package notification

import (
	entity "NFTM/shared/entities/notification"
	dynamodb_helper "NFTM/shared/repositories/dynamodb"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ListByDateArgs struct {
	UserID *string
	Limit  *int32
	From   *string
}

func ListByDate(ctx context.Context, args ListByDateArgs) (*[]entity.Notification, *string, error) {
	ns := GetNotificationService()
	newCond := expression.Key("userId").Equal(expression.Value(*args.UserID))
	expr, err := expression.NewBuilder().WithKeyCondition(newCond).Build()

	if err != nil {
		return nil, nil, err
	}

	orders := &[]entity.Notification{}
	input := dynamodb.QueryInput{
		TableName:                 &ns.TableName,
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		IndexName:                 &ns.UserDateIndex,
		ScanIndexForward:          aws.Bool(false),
	}
	fmt.Printf("input: %v\n", input)
	if args.Limit != nil && *args.Limit > 0 {
		input.Limit = args.Limit
	}
	// fmt.Printf("args.From: %v\n", args.From)
	if args.From != nil && *args.From != "" {
		startKey, err := dynamodb_helper.Deserialize(*args.From, entity.Notification{})
		if err != nil {
			return nil, nil, err
		}
		input.ExclusiveStartKey = *startKey

		fmt.Printf("input.ExclusiveStartKey: %v\n", input.ExclusiveStartKey)
	}

	output, err := ns.Client.Query(ctx, &input)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, nil, err
	}

	err = attributevalue.UnmarshalListOfMaps(output.Items, orders)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, nil, err
	}

	next, err := dynamodb_helper.Serialize(output.LastEvaluatedKey, &entity.Notification{})
	if err != nil {
		return nil, nil, err
	}

	return orders, next, nil
}
