package dynamodb

import (
	"NFTM/shared/config"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func SetNFTItemTable(stack constructs.Construct) awsdynamodb.Table {
	println("TableName:", config.Conf.Tables[config.NFT_ITEM].Name)

	return awsdynamodb.NewTable(stack, jsii.String(config.Conf.Tables[config.NFT_ITEM].Name), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("collectionAddress"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("tokenId"),
			Type: awsdynamodb.AttributeType_NUMBER,
		},
		TableName:   jsii.String(config.Conf.Tables[config.NFT_ITEM].Name),
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
	})
}
