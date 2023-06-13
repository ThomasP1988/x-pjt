package main

import (
	"aws/cognito"
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewAdminStack(scope constructs.Construct, id *string, props *AwsStackProps) (awscdk.Stack, cognito.AdminPoolResponse) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, id, &sprops)
	fmt.Printf("stack: %v\n", stack)

	userPool := cognito.SetAdminPool(stack, cognito.AdminProps{
		Stage: stage,
	})

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
