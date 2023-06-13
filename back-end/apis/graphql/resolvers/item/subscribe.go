package nftitem

import (
	"NFTM/apis/graphql/utils"
	"NFTM/shared/entities/nft"
	"context"
	"log"
)

func Subscribe(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {
	log.Printf("args: %v\n", args)
	return nft.Item{
		TokenID: 0,
	}, nil
}
