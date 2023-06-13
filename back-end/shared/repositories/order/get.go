package order

import (
	entity "NFTM/shared/entities/order"
	dynamodb_helper "NFTM/shared/repositories/dynamodb"
	"NFTM/shared/repositories/log"
	"context"
	"fmt"
)

func Get(ctx context.Context, orderID string) (*entity.Order, error) {
	os := GetOrderService()
	order := &entity.Order{}

	doesntExist, err := dynamodb_helper.GetOne(os.Client, &os.TableName, order, map[string]interface{}{
		"orderId": orderID,
	}, nil)

	if err != nil {
		log.Error("Error getting order", err)
		return nil, err
	}

	if doesntExist {
		fmt.Printf("orderID: %v\n", orderID)
		log.Error("Order doesn't exist", nil)
	}

	return order, nil

}
