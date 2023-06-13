package user

import (
	"NFTM/apis/graphql/utils"
	app_errors "NFTM/shared/common/errors"
	user_service "NFTM/shared/repositories/user"
	"context"
	"fmt"
	"strings"
)

func Me(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {

	user, err := user_service.Get(ctx, *args.UserID)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, app_errors.ErrInternalServer
	}

	if err == nil && user == nil {
		AMR := strings.Split(args.Event.Identity.CognitoIdentityAuthProvider, "\"")
		userIDFromPool := strings.Split(AMR[3], ":")[2]
		user, err = utils.GetUser(ctx, userIDFromPool, *args.UserID)
	}

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, app_errors.ErrInternalServer
	}
	fmt.Printf("user: %v\n", user)
	return *user, nil
}
