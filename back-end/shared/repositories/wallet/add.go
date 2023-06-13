package wallet

import (
	entity "NFTM/shared/entities/wallet"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Add(context context.Context, walletAsset *entity.WalletAsset) error {
	ws := GetWalletService()

	walletAsset.LastUpdate = time.Now()

	marshalledItem, err := attributevalue.MarshalMap(walletAsset)
	if err != nil {
		println("order service Create, error marshalling item", err.Error())
		return err
	}
	fmt.Printf("marshalledItem: %v\n", walletAsset)
	input := &dynamodb.PutItemInput{
		Item:      marshalledItem,
		TableName: &ws.TableName,
	}

	_, err = ws.Client.PutItem(context, input)

	if err != nil {
		fmt.Printf("Create order: %v\n", err)
		return err
	}

	return nil
}
