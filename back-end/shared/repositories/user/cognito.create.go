package user

import (
	"NFTM/shared/config"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (cg *CognitoClient) Create(ctx context.Context, email string) (*string, error) {
	output, err := cg.Client.AdminCreateUser(ctx, &cognitoidentityprovider.AdminCreateUserInput{
		Username:   aws.String(email),
		UserPoolId: &config.Conf.User.UserPool,
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	return output.User.Username, nil
}
