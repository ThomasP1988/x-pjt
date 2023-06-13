package nft

import (
	"NFTM/shared/config"
	"time"
)

type Collection struct {
	Address       string            `dynamodbav:"address" json:"id"`
	Symbol        string            `dynamodbav:"symbol" json:"symbol"`
	ChainSymbol   string            `dynamodbav:"chainSymbol" json:"chainSymbol"`
	Description   string            `dynamodbav:"description" json:"description,omitempty"`
	Name          string            `dynamodbav:"name" json:"name"`
	ChainName     string            `dynamodbav:"chainName" json:"chainName"`
	OpenseaSlug   string            `dynamodbav:"openseaSlug" json:"openseaSlug"`
	Status        CollectionStatus  `dynamodbav:"status" json:"status"`
	Supply        int               `dynamodbav:"supply" json:"supply"`
	Chain         config.Blockchain `dynamodbav:"chain" json:"chain"`
	ImagePath     string            `dynamodbav:"imagePath" json:"imagePath,omitempty"`
	ThumbnailPath string            `dynamodbav:"thumbnailPath" json:"thumbnailPath,omitempty"`
	FirstItemID   int               `dynamodbav:"firstItemId" json:"firstItemId"`
	SubmittedBy   string            `dynamodbav:"submittedBy" json:"submittedBy"`
	SubmittedAt   time.Time         `dynamodbav:"submittedAt" json:"submittedAt"`
	ValidatedBy   string            `dynamodbav:"validatedBy" json:"validatedBy"`
	ValidatedAt   time.Time         `dynamodbav:"validatedAt" json:"validatedAt"`
}
