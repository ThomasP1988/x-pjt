package blockchain

import (
	"NFTM/shared/config"
	"NFTM/shared/entities/nft"
	"context"
	"fmt"
	"testing"
)

func TestFetchItem(t *testing.T) {
	// "0xaadc2d4261199ce24a4b0a57370c4fcf43bb60aa", // THE CURRENCY
	// asset: "0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb", // CRYPTO PUNKS
	ctx := context.Background()

	collection := nft.Collection{
		Address: "0xd7AA15dFe0b67218Afe38EcbeE57AF8542D3315f",
		Chain:   config.Palm,
	}

	item, err := FetchItem(ctx, &collection, 1)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}

	fmt.Printf("item: %v\n", item.Attributes[0])

	t.Fail()

}
