package cognito

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func SetPostConfirmationLambda(stack constructs.Construct, props Props) awslambda.Function {
	fmt.Printf("\"SetPostConfirmationLambda\": %v\n", "SetPostConfirmationLambda")
	fmt.Printf("props.LambdaEnv: %v\n", props.LambdaEnv)
	fmt.Printf("props.LambdaRights: %v\n", props.LambdaRights)
	lambdaCognito := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("ddb-stream-GraphQL"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:         jsii.String("../../nftm/aws/cognito/post_confirmation"),
		InitialPolicy: &[]awsiam.PolicyStatement{*props.LambdaRights},
		Environment:   props.LambdaEnv,
	})

	return lambdaCognito
}
