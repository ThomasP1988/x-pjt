package helper

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func AWSPrint(stack awscdk.Stack, key string, value string, description *string) {
	awscdk.NewCfnOutput(stack, jsii.String(key), &awscdk.CfnOutputProps{
		Value:       jsii.String(value),
		Description: description,
		ExportName:  jsii.String(key),
	})
}
