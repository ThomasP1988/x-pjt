package dynamodb

import (
	"NFTM/shared/config"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func SetNotificationTable(stack constructs.Construct) awsdynamodb.Table {
	println("TableName:", config.Conf.Tables[config.NOTIFICATION].Name)

	table := awsdynamodb.NewTable(stack, jsii.String(config.Conf.Tables[config.NOTIFICATION].Name), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("userId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		Stream:      awsdynamodb.StreamViewType_NEW_IMAGE,
		TableName:   jsii.String(config.Conf.Tables[config.NOTIFICATION].Name),
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
	})

	table.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String(config.Conf.Tables[config.NOTIFICATION].SecondaryIndex[config.UserDateIndex]),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("userId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("createdAt"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})
	return table
}
