package dynamodb

import (
	"NFTM/shared/config"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func SetCollectionTable(stack constructs.Construct) {
	println("TableName:", config.Conf.Tables[config.COLLECTION].Name)

	table := awsdynamodb.NewTable(stack, jsii.String(config.Conf.Tables[config.COLLECTION].Name), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("address"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName:   jsii.String(config.Conf.Tables[config.COLLECTION].Name),
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
	})
	// fmt.Printf("table: %v\n", table)

	table.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String(config.Conf.Tables[config.COLLECTION].SecondaryIndex[config.SymbolIndex]),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("symbol"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	table.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String(config.Conf.Tables[config.COLLECTION].SecondaryIndex[config.StatusIndex]),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("status"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("submittedAt"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	table.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String(config.Conf.Tables[config.COLLECTION].SecondaryIndex[config.SubmittedAtIndex]),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("submittedAt"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

}
