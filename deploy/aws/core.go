package main

import (
	"aws/dynamodb"
	"aws/s3"
	"aws/sns"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

var SharedCoreResources *CoreResources

type CoreResources struct {
	CustomerBucket *awss3.Bucket
	NFTItemTable   awsdynamodb.Table
}

func NewCoreStack(scope constructs.Construct, id *string, props *AwsStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	resources := CoreResources{}
	stack := awscdk.NewStack(scope, id, &sprops)
	resources.CustomerBucket = s3.SetMediaBucket(stack)
	dynamodb.SetCollectionTable(stack)
	resources.NFTItemTable = dynamodb.SetNFTItemTable(stack)
	dynamodb.SetBlockchainWalletTable(stack)
	dynamodb.SetTokenTable(stack)

	sns.SetFetchItemSNS(stack, sns.Props{
		Stage:        stage,
		LambdaEnv:    LambdaEnv,
		LambdaRights: &LambdaRights,
	})

	// k8s.SetKubernetesCluster(stack)
	SharedCoreResources = &resources
	return stack
}
