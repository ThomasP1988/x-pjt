package search

import (
	"NFTM/apis/graphql/utils"
	"context"
	"log"

	app_errors "NFTM/shared/common/errors"
	collection_repo "NFTM/shared/repositories/collection"
)

const (
	From   string = "from"
	Limit  string = "limit"
	Text   string = "text"
	Status string = "status"
)

const (
	Results string = "results"
	Total   string = "total"
)

func SearchCollection(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {

	var from string
	var status string
	var limit int32

	if (*args.Args)[From] != nil {
		from = (*args.Args)[From].(string)
	}

	if (*args.Args)[Limit] != nil {
		limit = int32((*args.Args)[Limit].(int))
	}
	if (*args.Args)[Status] != nil {
		status = (*args.Args)[Status].(string)
	}

	log.Printf("from: %v\n", from)
	log.Printf("limit: %v\n", limit)
	log.Printf("status: %v\n", status)

	// TODO: proper elasticsearch query
	collections, next, err := collection_repo.List(ctx, collection_repo.ListArgs{
		Limit: &limit,
		From:  &from,
	})

	log.Printf("next: %v\n", next)

	if err != nil {
		log.Printf("collection_repo.List: %v\n", err)
		return nil, app_errors.ErrInternalServer
	}

	return map[string]interface{}{
		Total:   20,
		Results: ResolveArray(*collections),
	}, nil
}
