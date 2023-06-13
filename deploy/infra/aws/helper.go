package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	configAWS "github.com/aws/aws-sdk-go-v2/config"
)

func GetCredentials() (aws.Credentials, error) {
	cfg, err := configAWS.LoadDefaultConfig(context.TODO(), configAWS.WithSharedConfigProfile("nftquant"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	return cfg.Credentials.Retrieve(context.TODO())
}
