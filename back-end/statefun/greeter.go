package main

import (
	"fmt"

	"github.com/apache/flink-statefun/statefun-sdk-go/v3/pkg/statefun"
)

type Greeter struct {
	SeenCount statefun.ValueSpec
}

func (g *Greeter) Invoke(ctx statefun.Context, message statefun.Message) error {
	if !message.Is(statefun.StringType) {
		return fmt.Errorf("unexpected message type %s", message.ValueTypeName())
	}

	var name string
	_ = message.As(statefun.StringType, &name)

	storage := ctx.Storage()

	var count int32
	storage.Get(g.SeenCount, &count)

	count += 1

	storage.Set(g.SeenCount, count)

	ctx.Send(statefun.MessageBuilder{
		Target: statefun.Address{
			FunctionType: statefun.TypeNameFrom("com.example.fns/inbox"),
			Id:           name,
		},
		Value: fmt.Sprintf("Hello %s for the %dth time!", name, count),
	})

	return nil
}
