package token

import (
	"NFTM/apis/graphql/utils"
	token_entity "NFTM/shared/entities/token"
	token_repo "NFTM/shared/repositories/token"
	"context"
)

func CreateNonce(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {

	token := token_entity.CreateNonceToken(*args.UserID)
	token_repo.Create(ctx, token)
	return map[string]string{
		"token": token.TokenID,
		"nonce": token.Nonce,
	}, nil
}
