package collection

import (
	"NFTM/apis/graphql/utils"
	app_errors "NFTM/shared/common/errors"
	collection_repo "NFTM/shared/repositories/collection"
	"context"
	"fmt"
)

const (
	CollectionAddress string = "collectionAddress"
)

func GetItem(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {
	var collectionAddress string

	if (*args.Args)[CollectionAddress] != nil {
		collectionAddress = (*args.Args)[CollectionAddress].(string)
	}

	collection, err := collection_repo.Get(ctx, collectionAddress)

	if err != nil {
		fmt.Printf("err collection_repo.Get: %v\n", err)
		return nil, app_errors.ErrCollectionFetching
	}

	if collection == nil {
		return nil, app_errors.ErrCollectionNotFound
	}

	return collection, nil
}
