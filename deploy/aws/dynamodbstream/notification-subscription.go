package dynamodbstream

import (
	"NFTM/shared/config"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type NotificationSubscriptionProps struct {
	Stage        config.Stage
	LambdaEnv    *map[string]*string
	LambdaRights awsiam.PolicyStatement
}

func SetNotificationTableConsumer(stack constructs.Construct, sourceArn *string, props *NotificationSubscriptionProps) awslambda.Function {

	lambdaConsumer := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("ddb-stream-notification"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:         jsii.String("../../nftm/aws/dynamodb/notification"),
		InitialPolicy: &[]awsiam.PolicyStatement{props.LambdaRights},
		Environment:   props.LambdaEnv,
	})

	awslambda.NewEventSourceMapping(stack, jsii.String("SourceMappingDDBEventsStream"), &awslambda.EventSourceMappingProps{
		BatchSize:        jsii.Number(1),
		Enabled:          jsii.Bool(true),
		RetryAttempts:    jsii.Number(3),
		StartingPosition: awslambda.StartingPosition_TRIM_HORIZON,
		EventSourceArn:   sourceArn,
		Target:           lambdaConsumer,
	})
	return lambdaConsumer
}
