package nftitem

import (
	entity "NFTM/shared/entities/nft"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type KeysInput struct {
	CollectionAddress string
	TokenId           string
}

func ListSomeByKeys(ctx context.Context, keys *[]KeysInput) (*[]entity.Item, error) {

	nfts := GetNFTItemService()

	items := &[]entity.Item{}
	for i := 0; i < len(*keys); i += batchSize {

		batchIndex := i + batchSize

		if batchIndex > len(*keys) {
			batchIndex = len(*keys)
		}

		keysMap := GetKeysMap((*keys)[i:batchIndex])

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

func GetKeysMap(keys []KeysInput) []map[string]types.AttributeValue {
	keysMap := []map[string]types.AttributeValue{}
	for i := range keys {
		keysMap = append(keysMap, map[string]types.AttributeValue{
			"collectionAddress": &types.AttributeValueMemberS{
				Value: keys[i].CollectionAddress,
			},
			"tokenId": &types.AttributeValueMemberN{
				Value: keys[i].TokenId,
			},
		})
	}
	return keysMap
}
