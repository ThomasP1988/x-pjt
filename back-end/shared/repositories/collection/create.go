package collection

import (
	entity "NFTM/shared/entities/nft"
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Create(ctx context.Context, collection *entity.Collection) error {
	cs := GetCollectionService()

	inc := 0
	for {
		existingCollection, err := GetBySymbol(ctx, collection.Symbol)

		if err != nil {
			return errors.New("error trying to check if symbol already exists")
		}

		if existingCollection == nil {
			break
		} else {
			if existingCollection.Address == collection.Address {
				return nil
			}
			inc++
			collection.Symbol = collection.Symbol + fmt.Sprint(inc)
		}
	}

	marshalledItem, err := attributevalue.MarshalMap(collection)
	if err != nil {
		println("collection service Create, error marshalling item", err.Error())
		return err
	}
	fmt.Printf("marshalledItem: %v\n", marshalledItem)

	input := &dynamodb.PutItemInput{
		Item:      marshalledItem,
		TableName: &cs.TableName,
	}

	_, err = cs.Client.PutItem(ctx, input)

	if err != nil {
		return err
	}

	return nil
}
