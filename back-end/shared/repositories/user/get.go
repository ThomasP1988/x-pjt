package user

import (
	entity "NFTM/shared/entities/user"
	dynamodb_helper "NFTM/shared/repositories/dynamodb"
	"NFTM/shared/repositories/log"
	"context"
	"fmt"
)

func Get(ctx context.Context, userID string) (*entity.User, error) {
	us := GetUserService()
	user := &entity.User{}

	doesntExist, err := dynamodb_helper.GetOne(us.Client, &us.TableName, user, map[string]interface{}{
		"id": userID,
	}, nil)

	if err != nil {
		log.Error("Error getting user", err)
		return nil, err
	}

	if doesntExist {
		fmt.Printf("doesnt exist userID: %v\n", userID)
		return nil, nil
	}

	return user, nil

}
