package sns

import (
	"NFTM/shared/config"
	"aws/helper"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type Props struct {
	Stage        config.Stage
	LambdaEnv    *map[string]*string
	LambdaRights *awsiam.PolicyStatement
}

func SetFetchItemSNS(stack constructs.Construct, props Props) *awssns.Topic {
	topic := awssns.NewTopic(stack, helper.SetName("FetchItem"), &awssns.TopicProps{
		TopicName: &config.Conf.SNS.FetchItemTopic,
	})

	lambdaSNS := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("SNS_fetchitems"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:         jsii.String("../../nftm/aws/sns/fetch_items"),
		InitialPolicy: &[]awsiam.PolicyStatement{*props.LambdaRights},
	})

	topic.AddSubscription(awssnssubscriptions.NewLambdaSubscription(lambdaSNS, &awssnssubscriptions.LambdaSubscriptionProps{}))

	awscdk.NewCfnOutput(stack, jsii.String("FetchItemTopic/ARN"), &awscdk.CfnOutputProps{
		Value: topic.TopicArn(),
	})

	return &topic
}
