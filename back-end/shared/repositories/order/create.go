package order

import (
	"NFTM/shared/constants"
	entity "NFTM/shared/entities/order"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Create(context context.Context, order *entity.Order) error {
	os := GetOrderService()
	if order.Status == constants.Order_OPEN {
		order.IsOpen = 1
	} else {
		order.IsOpen = 0
	}

	order.UserIDSymbolIsOpen = FormatUserIDSymbolIsOpen(order.UserID, order.Symbol, order.IsOpen)

	marshalledItem, err := attributevalue.MarshalMap(order)
	if err != nil {
		println("order service Create, error marshalling item", err.Error())
		return err
	}
	fmt.Printf("marshalledItem: %v\n", order)
	input := &dynamodb.PutItemInput{
		Item:      marshalledItem,
		TableName: &os.TableName,
	}

	_, err = os.Client.PutItem(context, input)

	if err != nil {
		fmt.Printf("Create order: %v\n", err)
		return err
	}

	return nil
}
