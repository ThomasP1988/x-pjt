package socket

import "context"

type PublishFunction func(ctx context.Context, key string, payload interface{}) error

var Publish *PublishFunction
