package dynamodb

import (
	"NFTM/shared/config"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func SetWalletAssetTable(stack constructs.Construct) {
	println("TableName:", config.Conf.Tables[config.WALLET_ASSET].Name)

	awsdynamodb.NewTable(stack, jsii.String(config.Conf.Tables[config.WALLET_ASSET].Name), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("userId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("assetId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName:   jsii.String(config.Conf.Tables[config.WALLET_ASSET].Name),
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
	})
}
