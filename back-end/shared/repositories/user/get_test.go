package user

import (
	"context"
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {

	user, err := Get(context.Background(), "test")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("user: %v\n", user)

	t.Fail()
}
