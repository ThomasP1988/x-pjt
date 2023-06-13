package blockchain

import (
	"NFTM/shared/repositories/collection"
	"context"
	"fmt"
	"testing"
)

func TestFetchItems(t *testing.T) {
	// "0xaadc2d4261199ce24a4b0a57370c4fcf43bb60aa", // THE CURRENCY
	// asset: "0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb", // CRYPTO PUNKS
	ctx := context.Background()
	tenderCollect, err := collection.Get(ctx, "TENDER")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}

	err = FetchItems(ctx, tenderCollect)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}

	t.Fail()

}
