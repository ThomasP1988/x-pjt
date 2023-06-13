package event

import (
	"time"

	"github.com/google/uuid"
)

type TypePaymentEvent string

const (
	PayementEventType_Deposit           TypePaymentEvent = "deposit"
	PayementEventType_MarketOrder       TypePaymentEvent = "marketOrder"
	PayementEventType_LimitOrder        TypePaymentEvent = "limitOrder"
	PayementEventType_LimitOrderPartial TypePaymentEvent = "limitOrderPartial"
	PayementEventType_Withdrawal        TypePaymentEvent = "withdrawal"
)

type PaymentEvent struct {
	ID                  string           `dynamodbav:"eventId"`
	UserID              string           `dynamodbav:"userId"`
	Type                TypePaymentEvent `dynamodbav:"type"`
	AmountAdded         int64            `dynamodbav:"amountAdded,omitempty"`
	CurrencyAdded       string           `dynamodbav:"currencyAdded,omitempty"`
	AmountSubstracted   int64            `dynamodbav:"amountSubstracted,omitempty"`
	CurrencySubstracted string           `dynamodbav:"currencySubstracted,omitempty"`
	CreatedAt           time.Time        `dynamodbav:"createdAt"`
}

func NewPaymentEvent(event *PaymentEvent) *PaymentEvent {
	event.ID = uuid.NewString()
	return event
}
