package main

import (
	"aws/cognito"
	"aws/dynamodb"
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewUserStack(scope constructs.Construct, id *string, props *AwsStackProps) (awscdk.Stack, cognito.PoolResponse) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, id, &sprops)
	fmt.Printf("\"!11111\": %v\n", "!")
	userPool := cognito.SetPool(stack, cognito.Props{
		Stage:        stage,
		LambdaEnv:    LambdaEnv,
		LambdaRights: &LambdaRights,
	})
	fmt.Printf("\"!22222\": %v\n", "!22222")
	dynamodb.SetUserTable(stack)

	protectedObj := awsiam.NewPolicyStatement(
		&awsiam.PolicyStatementProps{
			Effect:  awsiam.Effect_ALLOW,
			Actions: jsii.Strings("s3:GetObject"),
			Resources: jsii.Strings(
				*(*SharedCoreResources.CustomerBucket).BucketArn() + "/protected/*",
			),
		},
	)

	(*userPool.AuthRole).AddToPolicy(protectedObj)
	(*userPool.UnauthRole).AddToPolicy(protectedObj)

	return stack, userPool
}
