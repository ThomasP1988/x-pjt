package utils

import (
	"NFTM/shared/entities/notification"
	"NFTM/shared/entities/user"
	notif_service "NFTM/shared/repositories/notification"
	user_service "NFTM/shared/repositories/user"
	"context"
	"fmt"
	"time"
)

// tech debt: in postconfirmation its impossible to get identityId

func GetUser(ctx context.Context, poolID string, identityID string) (*user.User, error) {

	var user *user.User
	// 1. get user by identityID

	user, err := user_service.Get(ctx, identityID)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	// 2. if no check if by userID exists
	user, err = user_service.Get(ctx, poolID)

	if err != nil {
		return user, err
	}

	// 3. if yes replace old one with identityID
	if user != nil {
		user.ID = identityID
		err = user_service.Create(ctx, user)
		if err != nil {
			return nil, err
		}
		err = user_service.Delete(ctx, poolID)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		err = notif_service.Create(ctx, notification.Notification{
			UserID:  identityID,
			Type:    notification.Welcome_NotificationType,
			Message: "Welcome to NFT Quant" + time.Now().String(),
		})
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		return user, nil
	}
	return nil, nil
}
