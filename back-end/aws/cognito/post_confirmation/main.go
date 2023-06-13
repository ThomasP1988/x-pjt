package main

import (
	"context"
	"fmt"
	"time"

	"NFTM/shared/config"
	user "NFTM/shared/entities/user"
	user_service "NFTM/shared/repositories/user"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
)

func PostConfirmation(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	fmt.Printf("PostConfirmation for user: %s\n", event.UserName)

	fmt.Printf("event.Request.UserAttributes: %v\n", event.Request.UserAttributes)

	println("create wallet")
	// newWallet := wallet.Wallet{
	// 	UserID:              event.UserName,
	// 	LastUpdate:          time.Now(),
	// 	Own:                 map[string]int64{},
	// 	Available:           map[string]int64{},
	// 	OwnLastUpdate:       map[string]int64{},
	// 	AvailableLastUpdate: map[string]int64{},
	// }

	// err := wallet_service.Create(ctx, newWallet)

	// if err != nil {
	// 	fmt.Printf("err wallet: %v\n", err)
	// 	return event, err
	// }

	println("create user")
	newUser := user.User{
		ID:                   event.UserName,
		Email:                event.Request.UserAttributes["email"],
		CreatedAt:            time.Now(),
		LastSeenNotification: time.Now(),
	}

	err := user_service.Create(ctx, &newUser)

	if err != nil {
		fmt.Printf("err user: %v\n", err)
		return event, err
	}

	return event, nil
}

func main() {
	config.GetConfig(nil)
	lambda.Start(PostConfirmation)
}
