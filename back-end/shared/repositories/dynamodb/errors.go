package dynamodb_helper

import "errors"

var ErrDeserializedUnknownField error = errors.New("unknown field in provided start key")
var ErrDeserializedUnknownTypeField error = errors.New("unknown type of field in provided start key")
var ErrDeserializedUnparsableBool error = errors.New("unparsable bool in provided start key")
