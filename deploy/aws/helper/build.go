package helper

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/jsii-runtime-go"
)

func SetGoFunctionPropsGeneric(path string, buildCommand string, handler string, env *map[string]*string, rights awsiam.PolicyStatement) *awslambda.FunctionProps {
	var environment map[string]*string = map[string]*string{
		"CGO_ENABLED": jsii.String("0"),
		"GOOS":        jsii.String("linux"),
		"GOARCH":      jsii.String("amd64"),
	}

	return &awslambda.FunctionProps{
		Code: awslambda.NewAssetCode(jsii.String(path), &awss3assets.AssetOptions{
			Bundling: &awscdk.BundlingOptions{
				Image:       awslambda.Runtime_GO_1_X().BundlingImage(),
				User:        jsii.String("root"),
				Environment: &environment,
				Command: &[]*string{
					jsii.String("bash"),
					jsii.String("-c"),
					jsii.String(buildCommand),
				},
			},
		}),
		Handler:       jsii.String(handler),
		Timeout:       awscdk.Duration_Seconds(jsii.Number(300)),
		Runtime:       awslambda.Runtime_GO_1_X(),
		Environment:   env,
		InitialPolicy: &[]awsiam.PolicyStatement{rights},
		LogRetention:  awslogs.RetentionDays_FIVE_DAYS,
	}
}
