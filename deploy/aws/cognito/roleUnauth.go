package cognito

import (
	"aws/helper"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func CreateUnauthRole(stack constructs.Construct, userPoolIdentity *awscognito.CfnIdentityPool) *awsiam.Role {

	federated := awsiam.NewFederatedPrincipal(
		jsii.String("cognito-identity.amazonaws.com"),
		&map[string]interface{}{
			"StringEquals": map[string]string{
				"cognito-identity.amazonaws.com:aud": *(*userPoolIdentity).Ref(),
			},
			"ForAnyValue:StringLike": map[string]string{
				"cognito-identity.amazonaws.com:amr": "unauthenticated",
			},
		},
		jsii.String("sts:AssumeRoleWithWebIdentity"),
	)

	role := awsiam.NewRole(stack, helper.SetName("unauth-nftm"), &awsiam.RoleProps{
		AssumedBy: federated,
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect:    awsiam.Effect_ALLOW,
		Actions:   jsii.Strings("mobileanalytics:PutEvents", "cognito-sync:*"),
		Resources: jsii.Strings("*"),
	}))

	return &role
}
