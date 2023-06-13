package dynamodb

import (
	"NFTM/shared/config"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func SetOrderTable(stack constructs.Construct) {
	println("TableName:", config.Conf.Tables[config.ORDER].Name)

	table := awsdynamodb.NewTable(stack, jsii.String(config.Conf.Tables[config.ORDER].Name), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("orderId"),
			Type: awsdynamodb.AttributeType_STRING,
		},

		TableName:   jsii.String(config.Conf.Tables[config.ORDER].Name),
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
	})

	table.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String(config.Conf.Tables[config.ORDER].SecondaryIndex[config.UserIndex]),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("userId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("symbol"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	table.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String(config.Conf.Tables[config.ORDER].SecondaryIndex[config.UserIdSymbolIsOpenIndex]),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("userIdSymbolIsOpen"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("lastModified"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	table.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String(config.Conf.Tables[config.ORDER].SecondaryIndex[config.SymbolIndex]),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("symbol"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("isOpen"),
			Type: awsdynamodb.AttributeType_NUMBER,
		},
	})

}
