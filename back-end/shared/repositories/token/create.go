package token

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Create(context context.Context, token interface{}) error {
	ts := GetTokenService()
	marshalledItem, err := attributevalue.MarshalMap(token)
	if err != nil {
		println("token service Create, error marshalling item", err.Error())
		return err
	}
	fmt.Printf("marshalledItem: %v\n", token)
	input := &dynamodb.PutItemInput{
		Item:      marshalledItem,
		TableName: &ts.TableName,
	}

	_, err = ts.Client.PutItem(context, input)

	if err != nil {
		fmt.Printf("Create order: %v\n", err)
		return err
	}

	return nil
}
