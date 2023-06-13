package collection

import (
	"NFTM/apis/graphql/utils"
	app_errors "NFTM/shared/common/errors"
	"NFTM/shared/entities/nft"
	blockchain_libs "NFTM/shared/libs/blockchain"
	collection_repo "NFTM/shared/repositories/collection"
	nftitem "NFTM/shared/repositories/nft-item"
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	SubmitArgsAddress     = "address"
	SubmitArgsDescription = "description"
)

func Submit(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {

	address := (*args.Args)[SubmitArgsAddress].(string)
	description := (*args.Args)[SubmitArgsDescription].(string)
	log.Printf("address: %v\n", address)
	log.Printf("description: %v\n", description)

	// 1. check if collection already exist
	existingCollection, err := collection_repo.Get(ctx, address)

	if err != nil {
		return nil, app_errors.ErrInternalServer
	}

	if existingCollection != nil {
		return nil, app_errors.ErrCollectionAlreadyExist
	}
	log.Printf("\"fetching collection\": %v\n", "fetching collection")
	// 2. if no, fetch collection data
	collectionResult, err := blockchain_libs.FetchCollection(ctx, address)

	if err != nil {
		return nil, err
	}

	collection := &nft.Collection{
		Address:     common.HexToAddress(address).String(),
		Symbol:      collectionResult.Symbol,
		ChainSymbol: collectionResult.Symbol,
		Name:        collectionResult.Name,
		ChainName:   collectionResult.Name,
		Supply:      collectionResult.Supply,
		Chain:       collectionResult.Chain,
		SubmittedBy: *args.UserID,
		SubmittedAt: time.Now(),
		Description: description,
		Status:      nft.PendingValidation_CollectionStatus,
	}
	log.Printf("startfetch first item")

	// check if symbol already exist in collection, if yes, add a number after and recheck
	err = collection_repo.Create(ctx, collection)

	if err != nil {
		log.Println("error creating collection", err)
		return nil, err
	}

	firstIndex := 1

	_, err = nftitem.FetchItem(ctx, &nftitem.FetchItemMessage{
		CollectionAddress:      collection.Address,
		IndexItem:              &firstIndex,
		ShouldUpdateCollection: true,
	})

	if err != nil {
		log.Println("error sending item to fetch to SNS", err)
	}

	return collection, nil
}
