package aws

import (
	"NFTM/shared/config"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	configAWS "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

var AWSConfig *aws.Config
var ddbClient *dynamodb.Client
var snsClient *sns.Client
var s3Client *s3.Client

func SetAWSConfig() error {
	var cfg aws.Config
	var err error
	var options []func(*configAWS.LoadOptions) error = []func(*configAWS.LoadOptions) error{}

	if config.Conf.Region != nil {
		options = append(options, configAWS.WithRegion(*config.Conf.Region))
	}

	if config.Conf.Profile != nil {
		_, err := configAWS.LoadSharedConfigProfile(context.TODO(), *config.Conf.Profile)
		if err == nil {
			fmt.Printf("\"Profile exists\": %v\n", *config.Conf.Profile)
			options = append(options, configAWS.WithSharedConfigProfile(*config.Conf.Profile))
		}
	}

	cfg, err = configAWS.LoadDefaultConfig(context.TODO(), options...)
	// fmt.Printf("cfg: %v\n", cfg)
	if err != nil {
		println("unable to load SDK config, %v", err)
	}

	AWSConfig = &cfg

	return err
}

func GetAWSConfig() (*aws.Config, error) {
	if AWSConfig == nil {
		err := SetAWSConfig()
		if err != nil {
			return nil, err
		}
	}
	return AWSConfig, nil
}

func GetDynamoDBClient() *dynamodb.Client {
	var err error
	if AWSConfig == nil {
		err = SetAWSConfig()
	}

	if err != nil {
		log.Fatalf(err.Error())
	}

	if ddbClient == nil {
		ddbClient = dynamodb.NewFromConfig(*AWSConfig)
	}

	return ddbClient
}

func GetS3Client() *s3.Client {
	var err error
	if AWSConfig == nil {
		err = SetAWSConfig()
	}

	if err != nil {
		log.Fatalf(err.Error())
	}

	if s3Client == nil {
		s3Client = s3.NewFromConfig(*AWSConfig)

	}

	return s3Client
}

func GetSNSClient() *sns.Client {
	var err error
	if AWSConfig == nil {
		err = SetAWSConfig()
	}

	if err != nil {
		log.Fatalf(err.Error())
	}

	if snsClient == nil {
		snsClient = sns.NewFromConfig(*AWSConfig)
	}

	return snsClient
}
