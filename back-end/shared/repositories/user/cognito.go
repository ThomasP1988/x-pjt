package user

import (
	commonAws "NFTM/shared/common/aws"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoClient struct {
	Client *cognitoidentityprovider.Client
}

func NewCognitoClient() (*CognitoClient, error) {
	config, err := commonAws.GetAWSConfig()
	if err != nil {
		return nil, err
	}

	return &CognitoClient{
		Client: cognitoidentityprovider.NewFromConfig(*config),
	}, nil
}
