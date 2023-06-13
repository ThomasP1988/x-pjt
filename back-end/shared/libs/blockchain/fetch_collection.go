package blockchain

import (
	"NFTM/shared/config"
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type FetchCollectionResult struct {
	Symbol string
	Name   string
	Chain  config.Blockchain
	Supply int
}

func FetchCollection(ctx context.Context, address string) (*FetchCollectionResult, error) {
	tokenAddress := common.HexToAddress(address)
	client, chain, supply, err := DetectChain(ctx, tokenAddress)
	if err != nil {
		fmt.Println("error setting client", err)
		return nil, err
	}

	fmt.Printf("chain: %v\n", chain)
	fmt.Printf("supply: %v\n", supply)
	fmt.Printf("client: %v\n", client)

	contract721, err := NewERC721(tokenAddress, client)
	if err != nil {
		fmt.Println("error setting contract", err)
		return nil, err
	}

	collectionResult, err := getCollectionMetadata(ctx, contract721, tokenAddress)
	if err != nil {
		fmt.Println("error getting collection metadata", err)
		return nil, err
	}
	collectionResult.Chain = *chain
	collectionResult.Supply = supply

	return collectionResult, nil
}

func getCollectionMetadata(ctx context.Context, contract721 *ERC721, tokenAddress common.Address) (*FetchCollectionResult, error) {

	symbol, err := contract721.ERC721Caller.Symbol(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		fmt.Println("error getting symbol", err)
		return nil, err
	}

	name, err := contract721.ERC721Caller.Name(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		fmt.Println("error getting name", err)
		return nil, err
	}

	collectionResult := &FetchCollectionResult{
		Symbol: symbol,
		Name:   name,
	}

	return collectionResult, nil
}

func DetectChain(ctx context.Context, tokenAddress common.Address) (*ethclient.Client, *config.Blockchain, int, error) {
	var client *ethclient.Client
	var chain config.Blockchain
	var highestSupply *big.Int

	for k, v := range config.GetConfig(nil).Blockchains {
		fmt.Printf("blockchain NodeURL: %v\n", v.NodeURL)
		newClient, err := ethclient.Dial(v.NodeURL)

		if err != nil {
			fmt.Println("Oops! error setting blockchain client", err)
			fmt.Printf("k: %v\n", k)
			continue
		}
		contract721, err := NewERC721(tokenAddress, newClient)
		if err != nil {
			fmt.Println("error setting contract", err)
			fmt.Printf("k: %v\n", k)
			continue
		}

		supply, err := contract721.TotalSupply(&bind.CallOpts{
			Context: ctx,
		})

		if err != nil {
			fmt.Println("error getting supply", err)
			fmt.Printf("k: %v\n", k)
			continue
		}

		if highestSupply == nil || highestSupply.Int64() < supply.Int64() {
			highestSupply = supply
			chain = k
			client = newClient
		}
	}

	if highestSupply == nil || highestSupply.Int64() == 0 {
		return nil, nil, 0, errors.New("couldn't detect chain of contract, please verify address")
	}

	return client, &chain, int(highestSupply.Int64()), nil
}
