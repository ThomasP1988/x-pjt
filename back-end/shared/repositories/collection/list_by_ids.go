package collection

import (
	entity "NFTM/shared/entities/nft"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const batchSize int = 25

func ListByIds(ctx context.Context, ids *[]string) (*[]entity.Collection, error) {

	cs := GetCollectionService()
	collections := &[]entity.Collection{}

	for i := 0; i < len(*ids); i += batchSize {
		batchIndex := i + batchSize

		if batchIndex > len(*ids) {
			batchIndex = len(*ids)
		}

		keysMap := GetKeysMap((*ids)[i:batchIndex])

		input := dynamodb.BatchGetItemInput{
			RequestItems: map[string]types.KeysAndAttributes{
				cs.TableName: {
					Keys: keysMap,
				},
			},
		}

		batch, err := cs.Client.BatchGetItem(ctx, &input)
		if err != nil {
			fmt.Printf("Client.BatchGetItem collection err: %v\n", err)
			return nil, err
		}

		// fmt.Printf("batch: %v\n", batch)

		batchCollections := &[]entity.Collection{}

		err = attributevalue.UnmarshalListOfMaps(batch.Responses[cs.TableName], batchCollections)
		if err != nil {
			return nil, err
		}

		newCollections := append(*collections, (*batchCollections)...)
		collections = &newCollections
	}

	return collections, nil
}

func GetKeysMap(ids []string) []map[string]types.AttributeValue {
	keysMap := []map[string]types.AttributeValue{}
	for i := range ids {
		keysMap = append(keysMap, map[string]types.AttributeValue{
			"address": &types.AttributeValueMemberS{
				Value: (ids)[i],
			},
		})
	}
	return keysMap
}
