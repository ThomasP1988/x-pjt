package dynamodb

import (
	"NFTM/shared/config"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func SetTokenTable(stack constructs.Construct) {
	tableInfo := config.Conf.Tables[config.TOKEN]

	awsdynamodb.NewTable(stack, jsii.String(tableInfo.Name), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("tokenId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName:           jsii.String(tableInfo.Name),
		BillingMode:         awsdynamodb.BillingMode_PAY_PER_REQUEST,
		TimeToLiveAttribute: jsii.String("ttl"),
		PointInTimeRecovery: jsii.Bool(false),
	})

}
