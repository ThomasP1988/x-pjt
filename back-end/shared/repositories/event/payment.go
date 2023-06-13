package event

import (
	commonAws "NFTM/shared/common/aws"
	"NFTM/shared/config"
	entity "NFTM/shared/entities/event"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var service *PaymentEventService

type PaymentEventService struct {
	Client    *dynamodb.Client
	TableName string
}

func GetPaymentEventService() *PaymentEventService {
	if service == nil {
		service = &PaymentEventService{
			Client:    commonAws.GetDynamoDBClient(),
			TableName: config.Conf.Tables[config.PAYMENT_EVENT].Name,
		}
	}

	return service
}

func Add(ctx context.Context, paymentEvent *entity.PaymentEvent) error {
	pes := GetPaymentEventService()
	marshalledItem, err := attributevalue.MarshalMap(paymentEvent)
	if err != nil {
		println("payment event service, error marshalling item", err.Error())
		return err
	}
	fmt.Printf("marshalledItem: %v\n", marshalledItem)
	input := &dynamodb.PutItemInput{
		Item:      marshalledItem,
		TableName: &pes.TableName,
	}

	_, err = pes.Client.PutItem(ctx, input)

	if err != nil {
		return err
	}

	return nil

}
