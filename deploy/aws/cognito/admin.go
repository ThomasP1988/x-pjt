package cognito

import (
	"NFTM/shared/config"
	"aws/helper"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
)

type AdminProps struct {
	Stage config.Stage
}

type AdminPoolResponse struct {
	Pool       *awscognito.UserPool
	AuthRole   *awsiam.Role
	UnauthRole *awsiam.Role
}

func SetAdminPool(stack awscdk.Stack, props AdminProps) AdminPoolResponse {

	userPool := awscognito.NewUserPool(stack, jsii.String("admin"), &awscognito.UserPoolProps{
		UserPoolName: helper.SetName("admin"),
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
		SelfSignUpEnabled: jsii.Bool(false),
	})
	webFrontUserPoolClient := awscognito.NewUserPoolClient(stack, jsii.String("admin-front"), &awscognito.UserPoolClientProps{
		UserPool:           userPool,
		UserPoolClientName: jsii.String(string(props.Stage) + "-admin-web"),
		AuthFlows: &awscognito.AuthFlow{
			UserSrp: jsii.Bool(false),
		},
		GenerateSecret: jsii.Bool(false),
	})

	userPool.AddDomain(jsii.String("admin-domain"), &awscognito.UserPoolDomainOptions{
		CognitoDomain: &awscognito.CognitoDomainOptions{
			DomainPrefix: jsii.String(config.Conf.Admin.Domain),
		},
	})

	userPoolIdentity := awscognito.NewCfnIdentityPool(stack, jsii.String("admin-identity"), &awscognito.CfnIdentityPoolProps{
		IdentityPoolName:               jsii.String(string(props.Stage) + "-admin"),
		AllowUnauthenticatedIdentities: jsii.Bool(true),
		CognitoIdentityProviders: []interface{}{
			map[string]string{
				"clientId":     *webFrontUserPoolClient.UserPoolClientId(),
				"providerName": *userPool.UserPoolProviderName(),
			},
		},
	})

	authRole := CreateAuthRole(stack, &userPoolIdentity)
	unAuthRole := CreateUnauthRole(stack, &userPoolIdentity)

	awscognito.NewCfnIdentityPoolRoleAttachment(stack, helper.SetName("admin-identity-roles"), &awscognito.CfnIdentityPoolRoleAttachmentProps{
		IdentityPoolId: userPoolIdentity.Ref(),
		Roles: map[string]string{
			"authenticated":   *(*authRole).RoleArn(),
			"unauthenticated": *(*unAuthRole).RoleArn(),
		},
	})

	helper.AWSPrint(stack, "admin-cognio-user-pool", *userPool.UserPoolId(), nil)
	helper.AWSPrint(stack, "admin-cognio-user-pool-arn", *userPool.UserPoolArn(), nil)
	helper.AWSPrint(stack, "admin-cognio-user-pool-web-client", *webFrontUserPoolClient.UserPoolClientId(), nil)
	helper.AWSPrint(stack, "admin-cognio-user-pool-identity", *userPoolIdentity.Ref(), nil)

	return AdminPoolResponse{
		Pool:       &userPool,
		AuthRole:   authRole,
		UnauthRole: unAuthRole,
	}

}
