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

func PartiallyFilled(ctx context.Context, orderID string, quantityFilled int64) (*entity.Order, error) {
	os := GetOrderService()
	var updateBuild expression.UpdateBuilder = expression.UpdateBuilder{}

	updateBuild = updateBuild.Set(expression.Name("status"), expression.Value(string(constants.Order_PARTIALLY_FILLED)))
	updateBuild = updateBuild.Set(expression.Name("quantity"), expression.Minus(expression.Name("quantity"), expression.Value(quantityFilled)))
	updateBuild = updateBuild.Set(expression.Name("filledQuantity"), expression.Plus(expression.Name("filledQuantity"), expression.Value(quantityFilled)))
	updateBuild = updateBuild.Set(expression.Name("partiallyFilled"), expression.Value(true))
	updateBuild = updateBuild.Set(expression.Name("lastModified"), expression.Value(time.Now().Format(time.RFC3339)))
	builder, err := expression.NewBuilder().WithUpdate(updateBuild).Build()

	if err != nil {
		fmt.Printf("PartiallyFilled: %v", err)
	}

	params := dynamodb.UpdateItemInput{

		Key: map[string]types.AttributeValue{
			"orderId": &types.AttributeValueMemberS{
				Value: orderID,
			},
		},
		TableName:                 &os.TableName,
		ReturnValues:              types.ReturnValueAllNew,
		ExpressionAttributeNames:  builder.Names(),
		ExpressionAttributeValues: builder.Values(),
		UpdateExpression:          builder.Update(),
	}

	output, err := os.Client.UpdateItem(ctx, &params)

	if err != nil {
		fmt.Printf("PartiallyFilled: %v\n", err)
		return nil, err
	}

	updatedOrder := &entity.Order{}
	attributevalue.UnmarshalMap(output.Attributes, updatedOrder)

	return updatedOrder, nil
}
