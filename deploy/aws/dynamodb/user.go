package dynamodb

import (
	"NFTM/shared/config"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func SetUserTable(stack constructs.Construct) {
	println("TableName:", config.Conf.Tables[config.USER].Name)

	table := awsdynamodb.NewTable(stack, jsii.String(config.Conf.Tables[config.USER].Name), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName:   jsii.String(config.Conf.Tables[config.USER].Name),
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
	})

	table.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String(config.Conf.Tables[config.USER].SecondaryIndex[config.EmailIndex]),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("email"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})
}
