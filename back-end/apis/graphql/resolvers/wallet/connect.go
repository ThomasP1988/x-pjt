package wallet

import (
	"NFTM/apis/graphql/utils"
	app_errors "NFTM/shared/common/errors"
	token_entity "NFTM/shared/entities/token"
	blockchain_lib "NFTM/shared/libs/blockchain"
	token_repo "NFTM/shared/repositories/token"
	"context"
	"fmt"
)

const (
	Signature string = "signature"
	TokenID   string = "tokenId"
)

func ConnectWallet(ctx context.Context, args utils.ResolverArgs) (interface{}, error) {

	var tokenId string
	var signature string

	if (*args.Args)[TokenID] != nil && (*args.Args)[Signature] != nil {
		tokenId = (*args.Args)[TokenID].(string)
		signature = (*args.Args)[Signature].(string)
	} else {
		return nil, app_errors.ErrMissingParameters
	}

	token := token_entity.TokenNonce{}

	err := token_repo.Use(ctx, tokenId, &token)

	if err != nil {
		return nil, err
	}

	if token.UserID != *args.UserID {
		return nil, app_errors.ErrUnauthorisedServer
	}
	// blockchain_lib.VerifySignatureFromJS(signature, token.Nonce)
	address, err := blockchain_lib.VerifySignatureFromJS(signature, token.Nonce)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("address: %v\n", *address)
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}
