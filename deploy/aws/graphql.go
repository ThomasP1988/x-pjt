package main

import (
	"NFTM/shared/config"
	"aws/cognito"
	"aws/helper"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

var schemaPathBase string = "../../nftm/apis/graphql/schema"
var mergedSchemaFilename string = "schema.graphql"
var (
	Mutation     string = "Mutation"
	Query               = "Query"
	Subscription        = "Subscription"
)

type GraphQLEndpoint struct {
	userPoolResponse      *cognito.PoolResponse
	userPoolAdminResponse *cognito.AdminPoolResponse
}

func NewGraphQLStack(scope constructs.Construct, id *string, props *AwsStackProps, graphQLprops *GraphQLEndpoint) (awscdk.Stack, awscdkappsyncalpha.GraphqlApi) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, id, &sprops)

	MergeGraphQLSchema()

	lambda := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("lambda-GraphQL"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:         jsii.String("../../nftm/apis/graphql"),
		InitialPolicy: &[]awsiam.PolicyStatement{LambdaRights},
		Environment:   LambdaEnv,
	})

	api := awscdkappsyncalpha.NewGraphqlApi(stack, jsii.String("graphql-api"), &awscdkappsyncalpha.GraphqlApiProps{
		Name:   jsii.String("GraphQL-API"),
		Schema: awscdkappsyncalpha.Schema_FromAsset(jsii.String(schemaPathBase + "/" + mergedSchemaFilename)),
		AuthorizationConfig: &awscdkappsyncalpha.AuthorizationConfig{
			DefaultAuthorization: &awscdkappsyncalpha.AuthorizationMode{
				AuthorizationType: awscdkappsyncalpha.AuthorizationType_IAM,
			},
			AdditionalAuthorizationModes: &[]*awscdkappsyncalpha.AuthorizationMode{
				{
					AuthorizationType: awscdkappsyncalpha.AuthorizationType_API_KEY,
				},
			},
		},
	})

	if stage == config.DEV {
		apiKeyProtected := awscdkappsyncalpha.NewGraphqlApi(stack, jsii.String("graphql-api-unprotected"), &awscdkappsyncalpha.GraphqlApiProps{
			Name:   jsii.String("GraphQL-API"),
			Schema: awscdkappsyncalpha.Schema_FromAsset(jsii.String(schemaPathBase + "/" + mergedSchemaFilename)),
			AuthorizationConfig: &awscdkappsyncalpha.AuthorizationConfig{
				DefaultAuthorization: &awscdkappsyncalpha.AuthorizationMode{
					AuthorizationType: awscdkappsyncalpha.AuthorizationType_API_KEY,
				},
			},
		})
		setResolvers(&apiKeyProtected, lambda)
		helper.AWSPrint(stack, "key-appsync-url", *apiKeyProtected.GraphqlUrl(), nil)
		helper.AWSPrint(stack, "key-appsync-key", *apiKeyProtected.ApiKey(), nil)
	}
	setRights(stack, &api, graphQLprops)
	setResolvers(&api, lambda)
	helper.AWSPrint(stack, "protected-appsync", *api.GraphqlUrl(), nil)
	return stack, api

}

func setRights(stack constructs.Construct, api *awscdkappsyncalpha.GraphqlApi, graphQLprops *GraphQLEndpoint) {
	region := *awscdk.Stack_Of(stack).Region()
	account := *awscdk.Stack_Of(stack).Account()

	ApiPolicy := awsiam.NewPolicyStatement(
		&awsiam.PolicyStatementProps{
			Effect:  awsiam.Effect_ALLOW,
			Actions: jsii.Strings("appsync:GraphQL"),
			Resources: jsii.Strings(
				"arn:aws:appsync:" + region + ":" + account + ":apis/" + *(*api).ApiId() + "/*",
			),
		},
	)

	authRole := *graphQLprops.userPoolResponse.AuthRole
	unauthRole := *graphQLprops.userPoolResponse.UnauthRole
	authRole.AddToPolicy(ApiPolicy)
	unauthRole.AddToPolicy(ApiPolicy)

	authAdminRole := *graphQLprops.userPoolAdminResponse.AuthRole
	unauthAdminRole := *graphQLprops.userPoolAdminResponse.UnauthRole
	authAdminRole.AddToPolicy(ApiPolicy)
	unauthAdminRole.AddToPolicy(ApiPolicy)
}

func setResolvers(api *awscdkappsyncalpha.GraphqlApi, lambda awslambda.IFunction) {
	lambdaSource := (*api).AddLambdaDataSource(jsii.String("graphql-go-function"), lambda, &awscdkappsyncalpha.DataSourceOptions{})

	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Query,
		FieldName: jsii.String("listCollections"),
	})
	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Query,
		FieldName: jsii.String("listNotifications"),
	})
	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Query,
		FieldName: jsii.String("me"),
	})
	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Query,
		FieldName: jsii.String("searchCollections"),
	})
	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Mutation,
		FieldName: jsii.String("connectWallet"),
	})
	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Mutation,
		FieldName: jsii.String("createToken"),
	})
	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Mutation,
		FieldName: jsii.String("inviteUser"),
	})
	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Mutation,
		FieldName: jsii.String("setLastSeenNotification"),
	})
	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Mutation,
		FieldName: jsii.String("submitCollection"),
	})
	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Mutation,
		FieldName: jsii.String("validateCollection"),
	})
	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Subscription,
		FieldName: jsii.String("notification"),
	})
	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Subscription,
		FieldName: jsii.String("onUpdatedItem"),
	})

	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Query,
		FieldName: jsii.String("getItem"),
	})

	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Query,
		FieldName: jsii.String("listCollectionsByIds"),
	})

	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Query,
		FieldName: jsii.String("listItemsByIds"),
	})

	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Query,
		FieldName: jsii.String("getCollection"),
	})

	lambdaSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  jsii.String("Collection"),
		FieldName: jsii.String("items"),
	})

	noneDataSource := (*api).AddNoneDataSource(jsii.String("real_time_data"), &awscdkappsyncalpha.DataSourceOptions{})
	noneDataSource.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Mutation,
		FieldName: jsii.String("notify"),
		RequestMappingTemplate: awscdkappsyncalpha.MappingTemplate_FromString(jsii.String(`{
			"version": "2018-05-29",
			"payload": $util.toJson($context.arguments.input)
		}`)),
		ResponseMappingTemplate: awscdkappsyncalpha.MappingTemplate_FromString(jsii.String(`$util.toJson($context.result)`)),
	})

	tableNFTItem := (*api).AddDynamoDbDataSource(jsii.String("nft-item"), SharedCoreResources.NFTItemTable, &awscdkappsyncalpha.DataSourceOptions{})
	tableNFTItem.CreateResolver(&awscdkappsyncalpha.BaseResolverProps{
		TypeName:  &Mutation,
		FieldName: jsii.String("updateItem"),
		RequestMappingTemplate: awscdkappsyncalpha.MappingTemplate_DynamoDbQuery(
			awscdkappsyncalpha.KeyCondition_Eq(
				jsii.String("collectionAddress"), jsii.String("collectionAddress"),
			).And(awscdkappsyncalpha.KeyCondition_Eq(
				jsii.String("tokenId"), jsii.String("tokenId"),
			)),
			nil,
		),
		ResponseMappingTemplate: awscdkappsyncalpha.MappingTemplate_DynamoDbResultItem(),
	})

}

func MergeGraphQLSchema() {
	files, err := ioutil.ReadDir(schemaPathBase)
	if err != nil {
		log.Fatal(err)
	}
	mergeFile, err := os.Create(schemaPathBase + "/" + mergedSchemaFilename)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		filePath := schemaPathBase + "/" + f.Name()
		extension := filepath.Ext(filePath)
		if extension == ".gql" {
			contentFile, err := os.ReadFile(filePath)
			if err != nil {
				log.Fatal(err)
			}

			_, err = mergeFile.Write(contentFile)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("extension: %v\n", extension)
	}
}
