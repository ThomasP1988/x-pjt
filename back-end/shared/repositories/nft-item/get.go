package nftitem

import (
	"NFTM/shared/entities/nft"
	dynamodb_helper "NFTM/shared/repositories/dynamodb"
	"context"
)

func Get(ctx context.Context, collectionAddress string, tokenId int32, item *nft.Item) (bool, error) {
	nfts := GetNFTItemService()
	return dynamodb_helper.GetOne(nfts.Client, &nfts.TableName, item, map[string]interface{}{
		"collectionAddress": collectionAddress,
		"tokenId":           tokenId,
	}, nil)
}
