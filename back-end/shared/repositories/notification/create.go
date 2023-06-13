package notification

import (
	entity "NFTM/shared/entities/notification"
	"NFTM/shared/repositories/socket"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

func Create(ctx context.Context, notification entity.Notification) error {
	ns := GetNotificationService()
	notification.ID = uuid.NewString()
	notification.CreatedAt = time.Now()
	notification.Read = false

	marshalledItem, err := attributevalue.MarshalMap(notification)
	if err != nil {
		println("notification service Create, error marshalling item", err.Error())
		return err
	}
	fmt.Printf("marshalledItem: %v\n", marshalledItem)

	input := &dynamodb.PutItemInput{
		Item:      marshalledItem,
		TableName: &ns.TableName,
	}

	_, err = ns.Client.PutItem(ctx, input)

	if err != nil {
		return err
	}

	if socket.Publish != nil {
		(*socket.Publish)(ctx, "notification"+notification.UserID, notification)
	}

	return nil
}
