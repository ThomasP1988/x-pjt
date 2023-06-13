package notification

import (
	"NFTM/shared/config"
	entity "NFTM/shared/entities/notification"
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

var userIdTest string = "userIdTest"

func setup() {
	config.GetConfig(nil)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestList(t *testing.T) {

	ctx := context.Background()

	userId := "myID"

	err := Create(ctx, entity.Notification{
		ID:        "testid2",
		Type:      "welcom",
		UserID:    userId,
		CreatedAt: time.Now(),
		Message:   "new notification",
		Read:      false,
	})

	err = Create(ctx, entity.Notification{
		ID:        "testid2",
		Type:      "welcom",
		UserID:    userId,
		CreatedAt: time.Now(),
		Message:   "new notification",
		Read:      false,
	})

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
	limit := int32(1)
	// from := "myID/2022-01-17T23:42:55+02:00"
	notifications, next, err := ListByDate(ctx, ListByDateArgs{
		UserID: &userId,
		Limit:  &limit,
		// From:   &from,
	})

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}

	fmt.Printf("next: %v\n", *next)

	fmt.Printf("notifications: %v\n", notifications)

	t.Fail()

}
