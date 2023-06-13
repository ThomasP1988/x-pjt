package auth

import (
	"NFTM/shared/entities/auth"
	"NFTM/shared/services/auth/cognito"
)

type UserService interface {
	Auth(token string) (*auth.UserAuth, error)
}

func GetUserService() UserService {
	return cognito.GetService()
}
