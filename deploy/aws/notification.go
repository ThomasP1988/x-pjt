package main

import (
	"aws/dynamodb"
	"aws/dynamodbstream"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewNotificationStack(scope constructs.Construct, id *string, props *AwsStackProps, notifProps *dynamodbstream.NotificationSubscriptionProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, id, &sprops)

	table := dynamodb.SetNotificationTable(stack)

	notifProps.LambdaRights = awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Resources: &[]*string{
			table.TableStreamArn(),
			table.TableArn(),
		},
		Effect: awsiam.Effect_ALLOW,
		Actions: jsii.Strings(
			"kinesis:DescribeStream",
			"kinesis:PutRecord",
			"kinesis:PutRecords",
			"kinesis:GetShardIterator",
			"kinesis:GetRecords",
			"kinesis:ListShards",
			"kinesis:DescribeStreamSummary",
			"kinesis:RegisterStreamConsumer",
			"dynamodb:*",
		),
	})

	dynamodbstream.SetNotificationTableConsumer(stack, table.TableStreamArn(), notifProps)

	return stack
}
