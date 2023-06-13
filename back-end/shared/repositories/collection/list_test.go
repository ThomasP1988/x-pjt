package collection

import (
	"NFTM/shared/config"
	"context"
	"fmt"
	"os"
	"testing"
)

func setup() {
	config.GetConfig(nil)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestList(t *testing.T) {

	ctx := context.Background()

	// limit := int32(1)
	// from := "Address+0xaaDc2D4261199ce24A4B0a57370c4FCf43BB60aa/SubmittedAt+2022-03-01T14:48:41.685128689Z"
	// collections, next, err := GetCollectionService().List(ctx, ListArgs{
	// 	// UserID: &userId,
	// 	Limit: &limit,
	// 	// From:   &from,
	// })
	// collections, next, err := ListBySubmittedAt(ctx, ListBySubmittedAtArgs{
	// 	// UserID: &userId,
	// 	// Status: &nft.PendingValidation_CollectionStatus,
	// 	Limit: &limit,
	// 	// From:  &from,
	// })
	collections, err := ListByIds(ctx, &[]string{"0xF7EDbffca810DA0Ba8e196Ba6e5cf228dB432D45", "0xd7AA15dFe0b67218Afe38EcbeE57AF8542D3315f"})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("collections: %v\n", collections)

	t.Fail()

}
