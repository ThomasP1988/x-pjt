package dynamodb

import (
	"NFTM/shared/config"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func SetPaymentEventTable(stack constructs.Construct) {
	println("TableName:", config.Conf.Tables[config.PAYMENT_EVENT].Name)

	table := awsdynamodb.NewTable(stack, jsii.String(config.Conf.Tables[config.PAYMENT_EVENT].Name), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("eventId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName:   jsii.String(config.Conf.Tables[config.PAYMENT_EVENT].Name),
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
	})

	table.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String(config.Conf.Tables[config.PAYMENT_EVENT].SecondaryIndex[config.UserIndex]),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("userId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("type"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})
}
