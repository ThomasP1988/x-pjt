package token

import (
	"time"

	"github.com/google/uuid"
)

type TokenNonce struct {
	TokenID string    `dynamodbav:"tokenId"`
	Nonce   string    `dynamodbav:"nonce"`
	Ttl     time.Time `dynamodbav:"ttl"`
	UserID  string    `dynamodbav:"userId"`
}

func CreateNonceToken(userID string) TokenNonce {
	return TokenNonce{
		UserID:  userID,
		TokenID: "tkn-" + uuid.NewString(),
		Nonce:   uuid.NewString(),
		Ttl:     time.Now().Add(time.Minute),
	}
}
