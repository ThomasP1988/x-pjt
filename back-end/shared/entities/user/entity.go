package user

import "time"

type User struct {
	ID                   string    `dynamodbav:"id" json:"id"`
	Email                string    `dynamodbav:"email" json:"email"`
	LastSeenNotification time.Time `dynamodbav:"lastSeenNotification" json:"lastSeenNotification"`
	CreatedAt            time.Time `dynamodbav:"createdAt" json:"createdAt"`
}
