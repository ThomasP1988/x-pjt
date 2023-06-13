package collection

import (
	"NFTM/shared/entities/nft"
	"context"
	"fmt"
	"testing"
)

func TestUpdate(t *testing.T) {

	ctx := context.Background()

	// from := "myID/2022-01-17T23:42:55+02:00"
	// collections, next, err := GetCollectionService().List(ctx, ListArgs{
	// 	// UserID: &userId,
	// 	Limit: &limit,
	// 	// From:   &from,
	// })
	collection, err := Update(ctx, "0xaaDc2D4261199ce24A4B0a57370c4FCf43BB60aa", &nft.Collection{
		ChainSymbol: "test",
		ChainName:   "test2",
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("collections: %v\n", collection)

	t.Fail()

}
