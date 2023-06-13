package nftitem

import (
	entity "NFTM/shared/entities/nft"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const batchSize int = 25

func ListSomeByIdsAndCollection(ctx context.Context, collectionAddress string, ids *[]string) (*[]entity.Item, error) {

	nfts := GetNFTItemService()

	items := &[]entity.Item{}
	for i := 0; i < len(*ids); i += batchSize {

		batchIndex := i + batchSize

		if batchIndex > len(*ids) {
			batchIndex = len(*ids)
		}

		keysMap := GetKeysMapWithCollection(collectionAddress, (*ids)[i:batchIndex])

		input := dynamodb.BatchGetItemInput{
			RequestItems: map[string]types.KeysAndAttributes{
				nfts.TableName: {
					Keys: keysMap,
				},
			},
		}

		batch, err := nfts.Client.BatchGetItem(ctx, &input)
		if err != nil {
			fmt.Printf("Client.BatchGetItem collection err: %v\n", err)
			return nil, err
		}

		if err != nil {
			return nil, err
		}
		batchItems := &[]entity.Item{}

		err = attributevalue.UnmarshalListOfMaps(batch.Responses[nfts.TableName], batchItems)
		if err != nil {
			return nil, err
		}

		newItems := append(*items, (*batchItems)...)
		items = &newItems

	}

	return items, nil
}

func GetKeysMapWithCollection(collectionAddress string, ids []string) []map[string]types.AttributeValue {
	keysMap := []map[string]types.AttributeValue{}
	for i := range ids {
		keysMap = append(keysMap, map[string]types.AttributeValue{
			"collectionAddress": &types.AttributeValueMemberS{
				Value: collectionAddress,
			},
			"tokenId": &types.AttributeValueMemberN{
				Value: ids[i],
			},
		})
	}
	return keysMap
}
