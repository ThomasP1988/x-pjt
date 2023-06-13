package nftitem

import (
	"NFTM/apis/graphql/utils"
	app_errors "NFTM/shared/common/errors"
	"NFTM/shared/entities/nft"
	collection_repo "NFTM/shared/repositories/collection"
	nft_repo "NFTM/shared/repositories/nft-item"
	"context"
	"fmt"
)

const (
	CollectionAddress string = "collectionAddress"
	TokenID           string = "tokenId"
)

func GetItem(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {
	var collectionAddress string
	var tokenId int32

	if (*args.Args)[CollectionAddress] != nil {
		collectionAddress = (*args.Args)[CollectionAddress].(string)
	}

	if (*args.Args)[TokenID] != nil {
		tokenId = int32((*args.Args)[TokenID].(float64))
	}

	collection, err := collection_repo.Get(ctx, collectionAddress)

	if err != nil {
		fmt.Printf("err collection_repo.Get: %v\n", err)
		return nil, app_errors.ErrCollectionFetching
	}

	if collection == nil {
		return nil, app_errors.ErrCollectionNotFound
	}

	nftItem := nft.Item{}

	doesntExist, err := nft_repo.Get(ctx, collectionAddress, tokenId, &nftItem)

	if err != nil {
		fmt.Printf("collectionAddress: %v\n", collectionAddress)
		fmt.Printf("tokenId: %v\n", tokenId)
		fmt.Printf("err nft_repo.Get: %v\n", err)
		return nil, app_errors.ErrNFTItemRetrieving
	}

	if !doesntExist {
		return nftItem, nil
	}

	// create new nft item with isFetching status

	nftItem.CollectionAddress = collectionAddress
	nftItem.TokenID = int(tokenId)
	nftItem.IsFetching = true

	err = nft_repo.Create(ctx, &nftItem)

	if err != nil {
		fmt.Printf("err nft_repo.Create: %v\n", err)
		return nil, app_errors.ErrNFTItemCreating
	}

	tokenIdInt := int(tokenId)
	_, err = nft_repo.FetchItem(ctx, &nft_repo.FetchItemMessage{
		CollectionAddress:      collectionAddress,
		TokenID:                &tokenIdInt,
		ShouldUpdateCollection: false,
	})

	if err != nil {
		fmt.Printf("err nft_repo.FetchItem: %v\n", err)
		return nil, app_errors.ErrNFTItemFetching
	}

	return nftItem, nil
}
