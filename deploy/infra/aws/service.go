package aws

import (
	"NFTM/shared/config"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	cdk8splus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2"
)

type Props struct {
	MarketService cdk8splus.Service
}

func SetServiceAccount(scope constructs.Construct, id string, props *Props) *cdk8splus.ServiceAccount {

	serviceAccount := cdk8splus.NewServiceAccount(scope, jsii.String("aws-service"), &cdk8splus.ServiceAccountProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Annotations: &map[string]*string{
				"eks.amazonaws.com/role-arn": jsii.String("arn:aws:iam::" + config.Conf.AWSAccount + ":role/<IAM_ROLE_NAME>"),
			},
		},
	})

	return &serviceAccount
}
