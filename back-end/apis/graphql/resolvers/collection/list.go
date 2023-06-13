package collection

import (
	"NFTM/apis/graphql/utils"
	app_errors "NFTM/shared/common/errors"
	"context"
	"fmt"

	entity "NFTM/shared/entities/nft"
	repo "NFTM/shared/repositories/collection"
)

const (
	From   string = "from"
	Limit  string = "limit"
	Status string = "status"
)

func List(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {

	var status entity.CollectionStatus
	var from string
	var limit int32

	if (*args.Args)[Status] != nil {
		status = (*args.Args)[Status].(entity.CollectionStatus)
	}

	if (*args.Args)[From] != nil {
		from = (*args.Args)[From].(string)
	}

	if (*args.Args)[Limit] != nil {
		limit = int32((*args.Args)[Limit].(int))
	}

	var collections *[]entity.Collection
	var next *string
	var err error

	if (*args.Args)[Status] != nil {
		collections, next, err = repo.ListByStatus(ctx, repo.ListByStatusArgs{
			Status: &status,
			From:   &from,
			Limit:  &limit,
		})
	} else {
		collections, next, err = repo.ListBySubmittedAt(ctx, repo.ListBySubmittedAtArgs{
			From:  &from,
			Limit: &limit,
		})
	}

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, app_errors.ErrInternalServer
	}

	return map[string]interface{}{
		"next":        next,
		"collections": *collections,
	}, nil
}
