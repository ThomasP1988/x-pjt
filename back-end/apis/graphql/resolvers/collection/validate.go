package collection

import (
	"NFTM/apis/graphql/utils"
	app_errors "NFTM/shared/common/errors"
	"NFTM/shared/entities/nft"
	collection_repo "NFTM/shared/repositories/collection"
	"context"
	"log"
)

const (
	ValidateArgsAddress     = "address"
	ValidateArgsDescription = "description"
	ValidateArgsSymbol      = "symbol"
	ValidateArgsName        = "name"
	ValidateArgsImage       = "image"
	ValidateArgsOpenseaSlug = "openseaSlug"
	ValidateArgsNewStatus   = "status"
)

func Validate(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {
	log.Printf("validate")

	updateCollection := nft.Collection{}

	if (*args.Args)[ValidateArgsAddress] == nil {
		return nil, app_errors.ErrEmptyAddress
	}
	if (*args.Args)[ValidateArgsDescription] != nil {
		updateCollection.Description = (*args.Args)[ValidateArgsDescription].(string)
	}
	if (*args.Args)[ValidateArgsSymbol] != nil {
		updateCollection.Symbol = (*args.Args)[ValidateArgsSymbol].(string)
	}
	if (*args.Args)[ValidateArgsName] != nil {
		updateCollection.Name = (*args.Args)[ValidateArgsName].(string)
	}
	if (*args.Args)[ValidateArgsImage] != nil {
		updateCollection.ImagePath = (*args.Args)[ValidateArgsImage].(string)
	}
	if (*args.Args)[ValidateArgsOpenseaSlug] != nil {
		updateCollection.OpenseaSlug = (*args.Args)[ValidateArgsOpenseaSlug].(string)
	}
	if (*args.Args)[ValidateArgsNewStatus] != nil {
		updateCollection.Status = (*args.Args)[ValidateArgsNewStatus].(nft.CollectionStatus)
	}

	address := (*args.Args)[ValidateArgsAddress].(string)

	return collection_repo.Update(ctx, address, &updateCollection)
}
