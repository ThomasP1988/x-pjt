package main

import (
	"NFTM/shared/config"
	"NFTM/shared/entities/nft"
	"NFTM/shared/libs/blockchain"
	collection_repo "NFTM/shared/repositories/collection"
	nftitem "NFTM/shared/repositories/nft-item"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var collections map[string]*nft.Collection = map[string]*nft.Collection{}

func main() {
	config.GetConfig(nil)

	lambda.Start(Handler)
}

func Handler(ctx context.Context, req events.SNSEvent) {

	for _, sr := range req.Records {
		msgJSON := sr.SNS.Message

		msg := nftitem.FetchItemMessage{}

		fmt.Printf("msgJSON: %v\n", msgJSON)
		err := json.Unmarshal([]byte(msgJSON), &msg)

		if err != nil {
			fmt.Printf("err unmarshal json: %v\n", err)
			continue
		}

		var collec *nft.Collection
		var ok bool

		if _, ok = collections[msg.CollectionAddress]; !ok {
			collec, err = collection_repo.Get(ctx, msg.CollectionAddress)
			if err != nil {
				fmt.Printf("err finding collection : %v\n", err)
				continue
			}
			collections[msg.CollectionAddress] = collec
			fmt.Printf("collec: %v\n", collec)
		}
		var item *nft.Item
		if msg.TokenID != nil {
			item, err = blockchain.FetchItem(ctx, collections[msg.CollectionAddress], *msg.TokenID)
		} else if msg.IndexItem != nil {
			item, err = blockchain.FetchItemByIndex(ctx, collections[msg.CollectionAddress], *msg.IndexItem)
		} else {
			continue
		}

		if err != nil {
			fmt.Printf("err fetching item: %v\n", err)
			continue
		}

		if msg.ShouldUpdateCollection {
			updatedCollection, err := collection_repo.Update(ctx, msg.CollectionAddress, &nft.Collection{
				ImagePath:     item.ImagePath,
				ThumbnailPath: item.ThumbnailPath,
				FirstItemID:   item.TokenID,
			})

			if err != nil {
				fmt.Printf("err updating collection: %v\n", err)
				continue
			}

			collections[updatedCollection.Address] = updatedCollection
		}
	}
}
