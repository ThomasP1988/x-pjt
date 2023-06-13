package user

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestLastSeen(t *testing.T) {

	user, err := SetLastSeenNotification(context.Background(), "eu-west-1:636c6aa7-b464-4b6b-987e-ad29773844c1", time.Now())
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("user: %v\n", user)

	t.Fail()
}
