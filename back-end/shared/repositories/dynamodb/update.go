package dynamodb_helper

import (
	"NFTM/shared/constants"
	"context"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UpdateArgs struct {
	Client    *dynamodb.Client
	Ctx       context.Context
	Key       *map[string]types.AttributeValue
	TableName *string
	Item      interface{}
}

func Update(args *UpdateArgs) (*dynamodb.UpdateItemOutput, error) {
	var updateBuild expression.UpdateBuilder = expression.UpdateBuilder{}
	v := reflect.Indirect(reflect.ValueOf(args.Item))
	typesCollection := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if !reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			nameField := typesCollection.Field(i).Tag.Get(constants.Tag_DynamoDB)
			fmt.Printf("nameField: %v\n", nameField)
			updateBuild = updateBuild.Set(expression.Name(nameField), expression.Value(field.Interface()))
		}
	}

	updateBuilt := expression.NewBuilder().WithUpdate(updateBuild)
	builder, err := updateBuilt.Build()

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return nil, err
	}

	return args.Client.UpdateItem(args.Ctx, &dynamodb.UpdateItemInput{
		Key:                       *args.Key,
		TableName:                 args.TableName,
		ReturnValues:              types.ReturnValueAllNew,
		ExpressionAttributeNames:  builder.Names(),
		ExpressionAttributeValues: builder.Values(),
		UpdateExpression:          builder.Update(),
	})
}
