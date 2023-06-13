package main

import (
	"NFTM/shared/config"
	"aws/dynamodbstream"
	"aws/helper"
	"errors"
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	// "github.com/aws/jsii-runtime-go"
)

type AwsStackProps struct {
	awscdk.StackProps
}

var stage config.Stage = config.DEV

var LambdaEnv *map[string]*string = &map[string]*string{
	"stage": jsii.String(string(stage)),
}
var LambdaRights awsiam.PolicyStatement

func main() {
	app := awscdk.NewApp(nil)

	configContext := app.Node().TryGetContext(jsii.String("stage"))
	if configContext != nil {
		stage = config.Stage(configContext.(string))
	}

	println("stage", stage)

	if stage != config.DEV && stage != config.STAGING && stage != config.PROD {
		panic(errors.New("unknown stage"))
	}
	helper.SetStage(stage)

	config.GetConfig(&stage)

	LambdaRights = awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: &[]*string{
			jsii.String("execute-api:*"),
			jsii.String("dynamodb:*"),
			jsii.String("logs:*"),
			jsii.String("s3:*"),
			jsii.String("cognito:*"),
			jsii.String("lambda:Invoke"),
			jsii.String("cognito-idp:*"),
			jsii.String("sns:*"),
		},
		Resources: &[]*string{
			jsii.String("*"),
		},
	})

	NewCoreStack(app, helper.SetName("NFTM-Core"), &AwsStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	_, userPoolResponse := NewUserStack(app, helper.SetName("NFTM"), &AwsStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	_, adminPoolResponse := NewAdminStack(app, helper.SetName("NFTM-Admin"), &AwsStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	NewMarketStack(app, helper.SetName("NFTM-Market"), &AwsStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	_, api := NewGraphQLStack(app, helper.SetName("NFTM-GraphQL-appsync"), &AwsStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	}, &GraphQLEndpoint{
		userPoolResponse:      &userPoolResponse,
		userPoolAdminResponse: &adminPoolResponse,
	})

	NewNotificationStack(app, helper.SetName("NFTM-Notification"),
		&AwsStackProps{
			awscdk.StackProps{
				Env: env(),
			},
		},
		&dynamodbstream.NotificationSubscriptionProps{
			Stage: stage,
			LambdaEnv: &map[string]*string{
				"APIURL": api.GraphqlUrl(),
				"APIKey": api.ApiKey(),
			},
		},
	)

	app.Synth(nil)
	fmt.Printf("\"la2\": %v\n", "la")
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------

	fmt.Printf("config.Conf.Region: %v\n", *config.Conf.Region)
	return &awscdk.Environment{
		Account: jsii.String(config.Conf.AWSAccount),
		Region:  jsii.String(*config.Conf.Region),
	}

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
