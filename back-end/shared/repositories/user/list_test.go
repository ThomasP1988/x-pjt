package user

import (
	"NFTM/shared/config"
	"context"
	"fmt"
	"os"
	"testing"
)

func setup() {
	config.GetConfig(nil)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestList(t *testing.T) {

	result, next, err := List(context.Background(), UserListArgs{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("next: %v\n", next)
	for key, user := range *result {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("user: %v\n", user.ID)
		fmt.Printf("user: %v\n", user.CreatedAt)
		fmt.Printf("user: %v\n", user.Email)
		fmt.Printf("user: %v\n", user.LastSeenNotification)
	}

	// fmt.Printf("id: %v\n", id)

	t.Fail()
}
