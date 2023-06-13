package main

import (
	"aws/dynamodb"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func NewMarketStack(scope constructs.Construct, id *string, props *AwsStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, id, &sprops)

	dynamodb.SetWalletAssetTable(stack)
	dynamodb.SetOrderTable(stack)
	dynamodb.SetPaymentEventTable(stack)

	return stack
}
