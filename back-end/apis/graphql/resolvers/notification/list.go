package notification

import (
	"NFTM/apis/graphql/utils"
	app_errors "NFTM/shared/common/errors"
	repo "NFTM/shared/repositories/notification"
	"context"
)

const (
	From  string = "from"
	Limit string = "limit"
)

func List(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {

	var from string
	var limit int32

	if (*args.Args)[From] != nil {
		from = (*args.Args)[From].(string)
	}

	if (*args.Args)[Limit] != nil {
		limit = int32((*args.Args)[Limit].(float64))
	}

	notifications, next, err := repo.ListByDate(ctx, repo.ListByDateArgs{
		UserID: args.UserID,
		From:   &from,
		Limit:  &limit,
	})

	if err != nil {
		return nil, app_errors.ErrInternalServer
	}

	return map[string]interface{}{
		"next":          next,
		"notifications": notifications,
	}, nil
}
