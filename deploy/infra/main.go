package main

import (
	"infra/apis/grpc"
	"infra/ingresses/api"
	"infra/statefun"

	"NFTM/shared/config"

	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func main() {

	config.GetConfig(nil)

	app := cdk8s.NewApp(nil)

	// DATABASE

	// API
	apiMartketService, _ := grpc.SetAPIMarket(app, "market-api", nil)

	// INGRESS
	// auth.SetVirtualServer(app, "ingress-auth", &auth.Props{
	// 	MarketService: apiMartketService,
	// })
	api.SetAPIIngress(app, "ingress-api", &api.APIIngressProps{
		MarketService: apiMartketService,
	})

	statefun.SetStateFun(app, "state-fun", nil)

	app.Synth()
}
