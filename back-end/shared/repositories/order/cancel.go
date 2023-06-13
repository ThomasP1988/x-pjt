package order

import (
	"NFTM/shared/constants"
	entity "NFTM/shared/entities/order"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Cancel(ctx context.Context, orderID string) (*entity.Order, error) {
	os := GetOrderService()
	existingOrder, err := Get(ctx, orderID)

	var isOpen int8 = 0

	var updateBuild expression.UpdateBuilder = expression.UpdateBuilder{}
	updateBuild = updateBuild.Set(expression.Name("status"), expression.Value(string(constants.Order_CANCELLED)))
	updateBuild = updateBuild.Set(expression.Name("isOpen"), expression.Value(isOpen))
	updateBuild = updateBuild.Set(expression.Name("quantity"), expression.Value(0))
	updateBuild = updateBuild.Set(expression.Name("lastModified"), expression.Value(time.Now().Format(time.RFC3339)))
	updateBuild = updateBuild.Set(
		expression.Name("userIdSymbolIsOpen"),
		expression.Value(FormatUserIDSymbolIsOpen(existingOrder.UserID, existingOrder.Symbol, isOpen)),
	)

	builder, err := expression.NewBuilder().WithUpdate(updateBuild).Build()

	if err != nil {
		fmt.Printf("Filled: %v", err)
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
		fmt.Printf("Filled: %v\n", err)
		return nil, err
	}

	updatedOrder := &entity.Order{}
	attributevalue.UnmarshalMap(output.Attributes, updatedOrder)

	return updatedOrder, nil
}
