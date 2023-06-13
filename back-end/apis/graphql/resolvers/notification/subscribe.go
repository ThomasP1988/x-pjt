package notification

import (
	"NFTM/apis/graphql/utils"
	app_errors "NFTM/shared/common/errors"
	entity "NFTM/shared/entities/notification"
	"context"
	"log"
)

func Subscribe(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {

	log.Printf("args: %v\n", args)

	if *args.UserID != (*args.Args)["userId"].(string) {
		return nil, app_errors.ErrUnauthorisedServer
	}

	return entity.Notification{
		Type:    entity.Silent_Subscribed_NotificationType,
		Message: "",
	}, nil
}
