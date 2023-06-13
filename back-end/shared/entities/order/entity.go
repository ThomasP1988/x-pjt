package order

import (
	"NFTM/shared/constants"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID                 string                `dynamodbav:"orderId"`
	Status             constants.StatusOrder `dynamodbav:"status"`
	IsOpen             int8                  `dynamodbav:"isOpen"`
	UserID             string                `dynamodbav:"userId"`
	Symbol             string                `dynamodbav:"symbol"`
	UserIDSymbolIsOpen string                `dynamodbav:"userIdSymbolIsOpen"`
	Side               constants.SideOrder   `dynamodbav:"side"`
	Type               constants.TypeOrder   `dynamodbav:"type"`
	CreatedAt          time.Time             `dynamodbav:"createdAt"`
	LastModified       time.Time             `dynamodbav:"lastModified"`
	Price              int64                 `dynamodbav:"price"`
	Paid               int64                 `dynamodbav:"paid"`
	AveragePrice       int64                 `dynamodbav:"averagePrice"`
	Quantity           int64                 `dynamodbav:"quantity"`
	FilledQuantity     int64                 `dynamodbav:"filledQuantity"`
	OriginalQuantity   int64                 `dynamodbav:"originalQuantity"`
	PartiallyFilled    bool                  `dynamodbav:"partiallyFilled"`
}

func NewOrder(args Order) *Order {

	var status constants.StatusOrder = constants.Order_OPEN

	if args.Type == constants.Order_LIMIT {
		status = constants.Order_OPEN
	}

	return &Order{
		ID:               uuid.NewString(),
		Status:           status,
		UserID:           args.UserID,
		Symbol:           args.Symbol,
		Side:             args.Side,
		CreatedAt:        time.Now(),
		Price:            args.Price,
		Quantity:         args.Quantity,
		OriginalQuantity: args.OriginalQuantity,
		Type:             args.Type,
		PartiallyFilled:  false,
		LastModified:     time.Now(),
	}
}
