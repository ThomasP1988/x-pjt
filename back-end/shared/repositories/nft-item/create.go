package nftitem

import (
	entity "NFTM/shared/entities/nft"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Create(ctx context.Context, collection *entity.Item) error {
	nfts := GetNFTItemService()
	marshalledItem, err := attributevalue.MarshalMap(collection)
	if err != nil {
		println("nft item service Create, error marshalling item", err.Error())
		return err
	}
	fmt.Printf("marshalledItem: %v\n", marshalledItem)

	input := &dynamodb.PutItemInput{
		Item:      marshalledItem,
		TableName: &nfts.TableName,
	}

	output, err := nfts.Client.PutItem(ctx, input)

	if err != nil {
		return err
	} else {
		fmt.Printf("output: %v\n", output)
	}

	return nil
}
