package nft

type Item struct {
	TokenID           int              `dynamodbav:"tokenId" json:"id"`
	TokenURI          string           `dynamodbav:"tokenURI" json:"tokenURI"`
	IsFetching        bool             `dynamodbav:"isFetching" json:"isFetching"`
	CollectionAddress string           `dynamodbav:"collectionAddress" json:"collectionAddress"`
	ImageData         string           `dynamodbav:"image_data,omitempty" json:"image_data,omitempty"`
	Name              string           `dynamodbav:"name" json:"name"`
	Image             string           `dynamodbav:"image" json:"image"`
	ImagePath         string           `dynamodbav:"imagePath,omitempty" json:"imagePath,omitempty"`
	ThumbnailPath     string           `dynamodbav:"thumbnailPath,omitempty" json:"thumbnailPath,omitempty"`
	Description       string           `dynamodbav:"description" json:"description"`
	ExternalUrl       string           `dynamodbav:"externalUrl,omitempty" json:"external_url,omitempty"`
	BackgroundColor   string           `dynamodbav:"backgroundColor,omitempty" json:"background_color,omitempty"`
	AnimationUrl      string           `dynamodbav:"animationUrl,omitempty" json:"animation_url,omitempty"`
	YoutubeUrl        string           `dynamodbav:"youtubeUrl,omitempty" json:"youtube_url,omitempty"`
	Attributes        []ItemAttributes `dynamodbav:"attributes,omitempty" json:"attributes,omitempty"`
}

type ItemAttributes struct {
	DisplayType string `dynamodbav:"display_type,omitempty" json:"display_type,omitempty"`
	TraitType   string `dynamodbav:"trait_type,omitempty" json:"trait_type,omitempty"`
	Value       string `dynamodbav:"value,omitempty" json:"value,omitempty"`
	Color       string `dynamodbav:"color,omitempty" json:"color,omitempty"`
}
