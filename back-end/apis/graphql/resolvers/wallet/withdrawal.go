package wallet

import (
	"NFTM/apis/graphql/utils"
	app_errors "NFTM/shared/common/errors"
	"NFTM/shared/entities/notification"
	repo "NFTM/shared/repositories/notification"
	"context"
	"fmt"
	"time"
)

func WithdrawalNFT(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {

	var tokenId string
	var signature string

	if (*args.Args)[TokenID] != nil {
		tokenId = (*args.Args)[TokenID].(string)
	}

	if (*args.Args)[Signature] != nil {
		signature = (*args.Args)[Signature].(string)
	}
	fmt.Printf("signature: %v\n", signature)
	fmt.Printf("tokenId: %v\n", tokenId)

	err := repo.Create(ctx, notification.Notification{
		UserID:  *args.UserID,
		Type:    notification.Welcome_NotificationType,
		Message: "Welcome to NFT Quant" + time.Now().String(),
	})

	if err != nil {
		return nil, app_errors.ErrInternalServer
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}
