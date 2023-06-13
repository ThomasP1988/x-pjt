package nftitem

import (
	"NFTM/apis/graphql/utils"
	app_errors "NFTM/shared/common/errors"
	"context"
	"fmt"

	repo "NFTM/shared/repositories/nft-item"
)

const (
	Ids string = "ids"
)

func ListByIds(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {

	var ids []string

	if (*args.Args)[Ids] != nil {
		for i := range (*args.Args)[Ids].([]interface{}) {
			ids = append(ids, (*args.Args)[Ids].([]interface{})[i].(string))
		}
	}

	collection := args.Event.Source
	fmt.Printf("collection: %v\n", collection)

	items, err := repo.ListSomeByIdsAndCollection(ctx, collection["id"].(string), &ids)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, app_errors.ErrInternalServer
	}

	return items, nil
}
