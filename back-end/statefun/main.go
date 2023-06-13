package main

import (
	"NFTM/shared/config"
	"net/http"

	"github.com/apache/flink-statefun/statefun-sdk-go/v3/pkg/statefun"
)

var (
	port string = ":8000"
)

func main() {
	config.GetConfig(nil)
	greeter := &Greeter{
		SeenCount: statefun.ValueSpec{
			Name:      "seen_count",
			ValueType: statefun.Int32Type,
		},
	}

	builder := statefun.StatefulFunctionsBuilder()
	_ = builder.WithSpec(statefun.StatefulFunctionSpec{
		FunctionType: statefun.TypeNameFrom("com.nftm/greeter"),
		States:       []statefun.ValueSpec{greeter.SeenCount},
		Function:     greeter,
	})

	http.Handle("/statefun", builder.AsHandler())
	_ = http.ListenAndServe(port, nil)

}
