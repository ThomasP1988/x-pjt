package collection

import (
	entity "NFTM/shared/entities/nft"
	dynamodb_helper "NFTM/shared/repositories/dynamodb"
	"NFTM/shared/repositories/log"
	"context"
	"fmt"
)

func Get(ctx context.Context, address string) (*entity.Collection, error) {
	cs := GetCollectionService()
	collection := &entity.Collection{}

	doesntExist, err := dynamodb_helper.GetOne(cs.Client, &cs.TableName, collection, map[string]interface{}{
		"address": address,
	}, nil)

	if err != nil {
		log.Error("Error getting collection", err)
		return nil, err
	}

	if doesntExist {
		fmt.Printf("doesnt exist symbol: %v\n", address)
		return nil, nil
	}

	return collection, nil

}
