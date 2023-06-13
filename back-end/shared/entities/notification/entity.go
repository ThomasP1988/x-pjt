package notification

import "time"

type Notification struct {
	ID        string           `dynamodbav:"id" json:"id"`
	UserID    string           `dynamodbav:"userId" json:"userId"`
	Type      NotificationType `dynamodbav:"type" json:"type"`
	CreatedAt time.Time        `dynamodbav:"createdAt" json:"createdAt"`
	Message   string           `dynamodbav:"message" json:"message"`
	Read      bool             `dynamodbav:"read" json:"read"`
}
