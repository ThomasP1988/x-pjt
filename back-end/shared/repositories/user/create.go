package user

import (
	"context"
	"fmt"

	entity "NFTM/shared/entities/user"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Create(context context.Context, user *entity.User) error {
	us := GetUserService()
	marshalledItem, err := attributevalue.MarshalMap(user)
	if err != nil {
		println("user service Create, error marshalling item", err.Error())
		return err
	}
	fmt.Printf("marshalledItem: %v\n", marshalledItem)
	input := &dynamodb.PutItemInput{
		Item:      marshalledItem,
		TableName: &us.TableName,
	}
	fmt.Printf("us.TableName: %v\n", us.TableName)

	output, err := us.Client.PutItem(context, input)

	if err != nil {
		fmt.Printf("Create user: %v\n", err)
		return err
	}
	fmt.Printf("output: %v\n", output)
	return nil
}
