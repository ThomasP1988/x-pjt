package k8s

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func SetKubernetesCluster(stack constructs.Construct) *awseks.Cluster {

	cluster := awseks.NewCluster(stack, jsii.String("nft"), &awseks.ClusterProps{
		ClusterName:         jsii.String("nft"),
		Version:             awseks.KubernetesVersion_V1_21(),
		OutputClusterName:   jsii.Bool(true),
		OutputConfigCommand: jsii.Bool(true),
	})

	// auth := cluster.AwsAuth()
	awscdk.NewCfnOutput(stack, jsii.String("OIDC-issuer"), &awscdk.CfnOutputProps{
		Value: cluster.ClusterOpenIdConnectIssuer(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("OIDC-endpoint-url"), &awscdk.CfnOutputProps{
		Value: cluster.ClusterOpenIdConnectIssuerUrl(),
	})
	// auth.
	return &cluster
}
