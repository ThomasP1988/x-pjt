package main

import (
	"NFTM/shared/config"
	nftitem "NFTM/shared/repositories/nft-item"
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	config.GetConfig(nil)
	ctx := context.Background()
	// "0xaaDc2D4261199ce24A4B0a57370c4FCf43BB60aa", // THE CURRENCY
	// "0xd7AA15dFe0b67218Afe38EcbeE57AF8542D3315f", // EMPRESSES

	tokenId := 21275

	msg, err := json.Marshal(nftitem.FetchItemMessage{
		CollectionAddress:      "0xF7EDbffca810DA0Ba8e196Ba6e5cf228dB432D45",
		TokenID:                &tokenId,
		ShouldUpdateCollection: false,
	})

	if err != nil {
		fmt.Printf("err marshalling msg: %v\n", err)
	}

	Handler(ctx, events.SNSEvent{
		Records: []events.SNSEventRecord{
			{
				SNS: events.SNSEntity{
					Message: string(msg),
				},
			},
		},
	})

	t.Fail()

}
