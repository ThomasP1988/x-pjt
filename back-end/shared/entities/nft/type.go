package nft

type CollectionStatus string

var (
	PendingValidation_CollectionStatus CollectionStatus = "PENDING_VALIDATION"
	Accepted_CollectionStatus          CollectionStatus = "ACCEPTED"
	Denied_CollectionStatus            CollectionStatus = "DENIED"
)
