package user

import (
	"NFTM/apis/graphql/utils"

	app_errors "NFTM/shared/common/errors"
	user_entity "NFTM/shared/entities/user"
	"NFTM/shared/repositories/user"
	user_service "NFTM/shared/repositories/user"
	"context"
	"log"
	"time"
)

const (
	InviteArgsEmail = "email"
)

func Invite(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {

	log.Printf("invite")
	cognitoClient, err := user.NewCognitoClient()
	if err != nil {
		log.Printf("user.NewCognitoClient err: %v\n", err)
		return nil, app_errors.ErrInternalServer
	}

	email := (*args.Args)[InviteArgsEmail].(string)
	userID, err := cognitoClient.Create(ctx, email)

	if err != nil {
		log.Printf("cognitoClient.Create err: %v\n", err)
		return nil, err
	}

	newUser := &user_entity.User{
		ID:                   *userID,
		Email:                email,
		LastSeenNotification: time.Now(),
		CreatedAt:            time.Now(),
	}

	err = user_service.Create(ctx, newUser)

	if err != nil {
		return nil, app_errors.ErrInternalServer
	}

	return map[string]bool{
		"success": true,
	}, nil
}
