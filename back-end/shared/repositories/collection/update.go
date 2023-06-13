package collection

import (
	entity "NFTM/shared/entities/nft"
	dynamodb_helper "NFTM/shared/repositories/dynamodb"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Update(ctx context.Context, address string, updateCollection *entity.Collection) (*entity.Collection, error) {
	cs := GetCollectionService()
	result, err := dynamodb_helper.Update(&dynamodb_helper.UpdateArgs{
		Client: cs.Client,
		Ctx:    ctx,
		Key: &map[string]types.AttributeValue{
			"address": &types.AttributeValueMemberS{
				Value: address,
			},
		},
		TableName: &cs.TableName,
		Item:      updateCollection,
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	output := &entity.Collection{}

	err = attributevalue.UnmarshalMap(result.Attributes, output)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	return output, nil
}
