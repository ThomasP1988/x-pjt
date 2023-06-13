package blockchain

import (
	"NFTM/shared/entities/nft"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func FetchItemByIndex(ctx context.Context, collection *nft.Collection, index int) (*nft.Item, error) {

	contract721, err := GetContractERC721(ctx, collection)
	if err != nil {
		fmt.Println("error setting contract", err)
		return nil, err
	}

	return FetchItemDetailsByIndex(ctx, contract721, collection, index)
}

func FetchItem(ctx context.Context, collection *nft.Collection, tokenId int) (*nft.Item, error) {

	contract721, err := GetContractERC721(ctx, collection)
	if err != nil {
		fmt.Println("error setting contract", err)
		return nil, err
	}

	return FetchItemDetails(ctx, contract721, collection, tokenId)
}

func GetContractERC721(ctx context.Context, collection *nft.Collection) (*ERC721, error) {
	client, err := getClient(collection.Chain)

	if err != nil {
		fmt.Println("Error setting client", err)
		return nil, err
	}
	tokenAddress := common.HexToAddress(collection.Address)

	return NewERC721(tokenAddress, client)
}
