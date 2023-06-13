package user

import (
	"NFTM/apis/graphql/utils"
	"context"
	"fmt"
	"time"

	app_errors "NFTM/shared/common/errors"
	user_service "NFTM/shared/repositories/user"
)

func SetLastSeenNotification(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {
	user, err := user_service.SetLastSeenNotification(ctx, *args.UserID, time.Now())
	fmt.Printf("user: %v\n", user)
	if err != nil {
		return nil, app_errors.ErrInternalServer
	}

	return user, nil
}
