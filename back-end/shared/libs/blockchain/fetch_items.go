package blockchain

import (
	"NFTM/shared/entities/nft"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func FetchItems(ctx context.Context, collection *nft.Collection) error {
	client, err := getClient(collection.Chain)

	if err != nil {
		fmt.Println("Error setting client", err)
		return err
	}
	tokenAddress := common.HexToAddress(collection.Address)

	contract721, err := NewERC721(tokenAddress, client)
	if err != nil {
		fmt.Println("error setting contract", err)
		return err
	}

	for i := 0; i < collection.Supply; i++ {
		FetchItemDetails(ctx, contract721, collection, i)
	}

	return nil
}
