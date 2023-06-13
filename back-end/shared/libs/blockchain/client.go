package blockchain

import (
	"NFTM/shared/config"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

var blockchainClient map[config.Blockchain]*ethclient.Client = map[config.Blockchain]*ethclient.Client{}

func getClient(chain config.Blockchain) (*ethclient.Client, error) {
	if blockchainClient[chain] != nil {
		return blockchainClient[chain], nil
	}

	nodeURL := (config.GetConfig(nil)).Blockchains[chain].NodeURL

	fmt.Printf("nodeURL: %v\n", nodeURL)

	client, err := ethclient.Dial(nodeURL)

	if err != nil {
		return nil, err
	}

	blockchainClient[chain] = client

	return client, err
}
