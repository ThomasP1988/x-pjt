package collection

import (
	entity "NFTM/shared/entities/nft"
	dynamodb_helper "NFTM/shared/repositories/dynamodb"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ListArgs struct {
	Limit *int32
	From  *string
}

func List(ctx context.Context, args ListArgs) (*[]entity.Collection, *string, error) {
	cs := GetCollectionService()
	input := dynamodb.ScanInput{
		TableName: &cs.TableName,
	}

	if args.Limit != nil && *args.Limit > 0 {
		input.Limit = args.Limit
	}

	if args.From != nil && *args.From != "" {
		startKey, err := dynamodb_helper.Deserialize(*args.From, entity.Collection{})

		if err != nil {
			return nil, nil, err
		}

		input.ExclusiveStartKey = *startKey
		fmt.Printf("input.ExclusiveStartKey: %v\n", input.ExclusiveStartKey)
	}

	output, err := cs.Client.Scan(ctx, &input)

	if err != nil {
		return nil, nil, err
	}
	collections := &[]entity.Collection{}

	err = attributevalue.UnmarshalListOfMaps(output.Items, collections)
	if err != nil {
		return nil, nil, err
	}

	next, err := dynamodb_helper.Serialize(output.LastEvaluatedKey, &entity.Collection{})
	if err != nil {
		return nil, nil, err
	}

	return collections, next, nil

}
