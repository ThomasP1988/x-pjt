package collection

import (
	entity "NFTM/shared/entities/nft"
	dynamodb_helper "NFTM/shared/repositories/dynamodb"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ListByStatusArgs struct {
	Status *entity.CollectionStatus
	Limit  *int32
	From   *string
}

func ListByStatus(ctx context.Context, args ListByStatusArgs) (*[]entity.Collection, *string, error) {
	cs := GetCollectionService()
	newCond := expression.Key("status").Equal(expression.Value(*args.Status))
	expr, err := expression.NewBuilder().WithKeyCondition(newCond).Build()

	if err != nil {
		return nil, nil, err
	}

	collections := &[]entity.Collection{}
	input := dynamodb.QueryInput{
		TableName:                 &cs.TableName,
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		IndexName:                 &cs.StatusIndex,
		ScanIndexForward:          aws.Bool(false),
	}
	fmt.Printf("input: %v\n", input)
	if args.Limit != nil && *args.Limit > 0 {
		input.Limit = args.Limit
	}
	// fmt.Printf("args.From: %v\n", args.From)
	if args.From != nil && *args.From != "" {
		startKey, err := dynamodb_helper.Deserialize(*args.From, entity.Collection{})

		if err != nil {
			return nil, nil, err
		}

		input.ExclusiveStartKey = *startKey
		fmt.Printf("input.ExclusiveStartKey: %v\n", input.ExclusiveStartKey)
	}

	output, err := cs.Client.Query(ctx, &input)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, nil, err
	}

	err = attributevalue.UnmarshalListOfMaps(output.Items, collections)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, nil, err
	}
	fmt.Printf("output.LastEvaluatedKey: %v\n", output.LastEvaluatedKey)
	next, err := dynamodb_helper.Serialize(output.LastEvaluatedKey, &entity.Collection{})
	if err != nil {
		return nil, nil, err
	}

	return collections, next, nil
}
