package cognito

import (
	"NFTM/shared/config"
	"aws/helper"
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/customresources"
	"github.com/aws/jsii-runtime-go"
)

type Props struct {
	Stage        config.Stage
	LambdaEnv    *map[string]*string
	LambdaRights *awsiam.PolicyStatement
}

type PoolResponse struct {
	Pool       *awscognito.UserPool
	AuthRole   *awsiam.Role
	UnauthRole *awsiam.Role
}

func SetPool(stack awscdk.Stack, props Props) PoolResponse {
	fmt.Printf("SetPool props: %v\n", props)
	postConfirmationFunction := SetPostConfirmationLambda(stack, props)

	userPool := awscognito.NewUserPool(stack, jsii.String("customer"), &awscognito.UserPoolProps{
		UserPoolName: helper.SetName("customer"),
		SignInAliases: &awscognito.SignInAliases{
			Email: jsii.Bool(true),
		},
		PasswordPolicy: &awscognito.PasswordPolicy{
			MinLength:        jsii.Number(6),
			RequireDigits:    jsii.Bool(true),
			RequireLowercase: jsii.Bool(false),
			RequireSymbols:   jsii.Bool(false),
			RequireUppercase: jsii.Bool(false),
		},
		UserVerification: &awscognito.UserVerificationConfig{
			EmailStyle: awscognito.VerificationEmailStyle_CODE,
		},
		SelfSignUpEnabled: jsii.Bool(true),
		LambdaTriggers: &awscognito.UserPoolTriggers{
			PostConfirmation: postConfirmationFunction,
		},
	})

	openIdUserPoolClient := awscognito.NewUserPoolClient(stack, jsii.String("openId"), &awscognito.UserPoolClientProps{
		UserPool:           userPool,
		UserPoolClientName: jsii.String(string(props.Stage) + "-openId"),
		AuthFlows: &awscognito.AuthFlow{
			UserSrp: jsii.Bool(true),
		},
		GenerateSecret:       jsii.Bool(true),
		RefreshTokenValidity: awscdk.Duration_Days(jsii.Number(30)),
		OAuth: &awscognito.OAuthSettings{
			Scopes: &[]awscognito.OAuthScope{
				awscognito.OAuthScope_OPENID(),
				awscognito.OAuthScope_COGNITO_ADMIN(),
				awscognito.OAuthScope_PROFILE(),
			},
			Flows: &awscognito.OAuthFlows{
				AuthorizationCodeGrant: jsii.Bool(true),
			},
		},
	})

	webFrontUserPoolClient := awscognito.NewUserPoolClient(stack, jsii.String("web-front"), &awscognito.UserPoolClientProps{
		UserPool:           userPool,
		UserPoolClientName: jsii.String(string(props.Stage) + "-web"),
		AuthFlows: &awscognito.AuthFlow{
			UserSrp: jsii.Bool(false),
		},
		GenerateSecret: jsii.Bool(false),
	})

	userPool.AddDomain(jsii.String("customer-domain"), &awscognito.UserPoolDomainOptions{
		CognitoDomain: &awscognito.CognitoDomainOptions{
			DomainPrefix: jsii.String(config.Conf.User.Domain),
		},
	})

	userPoolIdentity := awscognito.NewCfnIdentityPool(stack, jsii.String("customer-identity"), &awscognito.CfnIdentityPoolProps{
		IdentityPoolName:               jsii.String(string(props.Stage) + "-customer"),
		AllowUnauthenticatedIdentities: jsii.Bool(true),
		CognitoIdentityProviders: []interface{}{
			map[string]string{
				"clientId":     *openIdUserPoolClient.UserPoolClientId(),
				"providerName": *userPool.UserPoolProviderName(),
			},
			map[string]string{
				"clientId":     *webFrontUserPoolClient.UserPoolClientId(),
				"providerName": *userPool.UserPoolProviderName(),
			},
		},
	})

	authRole := CreateAuthRole(stack, &userPoolIdentity)
	unAuthRole := CreateUnauthRole(stack, &userPoolIdentity)

	awscognito.NewCfnIdentityPoolRoleAttachment(stack, helper.SetName("customer-identity-roles"), &awscognito.CfnIdentityPoolRoleAttachmentProps{
		IdentityPoolId: userPoolIdentity.Ref(),
		Roles: map[string]string{
			"authenticated":   *(*authRole).RoleArn(),
			"unauthenticated": *(*unAuthRole).RoleArn(),
		},
	})

	helper.AWSPrint(stack, "cognio-user-pool", *userPool.UserPoolId(), nil)
	helper.AWSPrint(stack, "cognio-user-pool-arn", *userPool.UserPoolArn(), nil)
	helper.AWSPrint(stack, "cognio-user-pool-web-client", *webFrontUserPoolClient.UserPoolClientId(), nil)
	helper.AWSPrint(stack, "cognio-user-pool-openid-client", *openIdUserPoolClient.UserPoolClientId(), nil)
	helper.AWSPrint(stack, "cognio-user-pool-identity", *userPoolIdentity.Ref(), nil)

	PrintClientSecret(stack, openIdUserPoolClient, userPool)

	return PoolResponse{
		Pool:       &userPool,
		AuthRole:   authRole,
		UnauthRole: unAuthRole,
	}
}

func PrintClientSecret(stack awscdk.Stack, userPoolClient awscognito.UserPoolClient, userPool awscognito.UserPool) {
	describeCognitoUserPoolClient := customresources.NewAwsCustomResource(stack, jsii.String("userpool-description"), &customresources.AwsCustomResourceProps{
		ResourceType: jsii.String("Custom::DescribeCognitoUserPoolClient"),
		OnCreate: &customresources.AwsSdkCall{
			Region:  jsii.String(*config.Conf.Region),
			Service: jsii.String("CognitoIdentityServiceProvider"),
			Action:  jsii.String("describeUserPoolClient"),
			Parameters: map[string]string{
				"ClientId":   *userPoolClient.UserPoolClientId(),
				"UserPoolId": *userPool.UserPoolId(),
			},
			PhysicalResourceId: customresources.PhysicalResourceId_Of(userPoolClient.UserPoolClientId()),
		},
		Policy: customresources.AwsCustomResourcePolicy_FromSdkCalls(
			&customresources.SdkCallsPolicyOptions{
				Resources: customresources.AwsCustomResourcePolicy_ANY_RESOURCE(),
			},
		),
	})

	userPoolClientSecret := describeCognitoUserPoolClient.GetResponseField(jsii.String("UserPoolClient.ClientSecret"))

	helper.AWSPrint(stack, "cognio-user-client-secret", *userPoolClientSecret, nil)

}
