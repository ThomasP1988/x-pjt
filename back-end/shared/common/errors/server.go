package app_errors

import "errors"

var ErrInternalServer error = errors.New("internal server error")
var ErrUnauthorisedServer error = errors.New("unauthorised")
var ErrMissingParameters error = errors.New("missing parameters")

var ErrCollectionAlreadyExist error = errors.New("collection already exists")
var ErrCollectionFetching error = errors.New("error fetching collection")
var ErrCollectionNotFound error = errors.New("collection not found")

var ErrNFTItemRetrieving error = errors.New("error fetching nft")
var ErrNFTItemFetching error = errors.New("error fetching nft item details")
var ErrNFTItemCreating error = errors.New("error creating nft item")

var ErrEmptyAddress error = errors.New("empty address")
