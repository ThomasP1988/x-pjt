package order

import (
	entity "NFTM/shared/entities/order"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func UpdateQuantity(ctx context.Context, orderID string, newQuantity string) (*entity.Order, error) {
	os := GetOrderService()
	var updateBuild expression.UpdateBuilder = expression.UpdateBuilder{}
	updateBuild = updateBuild.Set(expression.Name("quantity"), expression.Value(newQuantity))
	updateBuild = updateBuild.Set(expression.Name("lastModified"), expression.Value(time.Now().Format(time.RFC3339)))

	builder, err := expression.NewBuilder().WithUpdate(updateBuild).Build()

	if err != nil {
		fmt.Printf("UpdateQuantity: %v", err)
	}

	output, err := os.Client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"orderId": &types.AttributeValueMemberS{
				Value: orderID,
			},
		},
		TableName:                 &os.TableName,
		ExpressionAttributeNames:  builder.Names(),
		ExpressionAttributeValues: builder.Values(),
		UpdateExpression:          builder.Update(),
		ReturnValues:              types.ReturnValueAllNew,
	})

	if err != nil {
		return nil, err
	}

	updatedOrder := &entity.Order{}
	attributevalue.UnmarshalMap(output.Attributes, updatedOrder)

	return updatedOrder, nil
}
