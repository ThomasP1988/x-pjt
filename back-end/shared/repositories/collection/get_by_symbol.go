package collection

import (
	entity "NFTM/shared/entities/nft"
	dynamodb_helper "NFTM/shared/repositories/dynamodb"
	"NFTM/shared/repositories/log"
	"context"
	"fmt"
)

func GetBySymbol(ctx context.Context, symbol string) (*entity.Collection, error) {
	cs := GetCollectionService()
	collection := &entity.Collection{}

	doesntExist, err := dynamodb_helper.GetOne(cs.Client, &cs.TableName, collection, map[string]interface{}{
		"symbol": symbol,
	}, &cs.SymbolIndex)

	if err != nil {
		log.Error("Error getting collection", err)
		return nil, err
	}

	if doesntExist {
		fmt.Printf("collection doesnt exist: %v\n", symbol)
		return nil, nil
	}

	return collection, nil
}
