package blockchain

import (
	"NFTM/shared/entities/nft"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func WithdrawalNFT(ctx context.Context, collection *nft.Collection, tokenID string) error {

	client, err := getClient(collection.Chain)

	if err != nil {
		fmt.Println("Error setting client", err)
		return err
	}

	tokenAddress := common.HexToAddress(collection.Address)

	contract721, err := NewERC721(tokenAddress, client)
	fmt.Printf("contract721: %v\n", contract721)
	// contract721.Approve(&bind.TransactOpts{}, "", 5)

	return nil

}
