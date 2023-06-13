package nftitem

import (
	"NFTM/apis/graphql/utils"
	app_errors "NFTM/shared/common/errors"
	"context"
	"fmt"

	repo "NFTM/shared/repositories/nft-item"
)

const (
	Keys string = "keys"
)

func ListByKeys(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {

	var keys []repo.KeysInput

	if (*args.Args)[Keys] != nil {
		keysInput := (*args.Args)[Keys].([]interface{})
		for i := range keysInput {
			keys = append(keys, repo.KeysInput{
				CollectionAddress: keysInput[i].(map[string]interface{})["collectionAddress"].(string),
				TokenId:           keysInput[i].(map[string]interface{})["tokenId"].(string),
			})
		}

	}

	items, err := repo.ListSomeByKeys(ctx, &keys)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, app_errors.ErrInternalServer
	}

	return items, nil
}
