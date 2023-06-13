package api

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	cdk8splus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2"
)

type APIIngressProps struct {
	cdk8s.ChartProps
	MarketService cdk8splus.Service
}

func SetAPIIngress(scope constructs.Construct, id string, props *APIIngressProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	cdk8splus.NewIngress(chart, jsii.String("ingress-api"), &cdk8splus.IngressProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Annotations: &map[string]*string{
				"kubernetes.io/ingress.class": jsii.String("nginx"),
				// "nginx.ingress.kubernetes.io/enable-cors":            jsii.String("true"),
				"nginx.ingress.kubernetes.io/rewrite-target": jsii.String("/$2"),
				// "nginx.ingress.kubernetes.io/cors-allow-methods":     jsii.String("PUT, GET, POST, OPTIONS"),
				// "nginx.ingress.kubernetes.io/cors-allow-credentials": jsii.String("true"),
				// "nginx.ingress.kubernetes.io/proxy-buffer-size":      jsii.String("256k"),
				// "nginx.ingress.kubernetes.io/backend-protocol":       jsii.String("http"),
				"nginx.ingress.kubernetes.io/websocket-services": props.MarketService.Name(),
				"nginx.org/websocket-services":                   props.MarketService.Name(),
				// "nginx.ingress.kubernetes.io/cors-allow-origin":      jsii.String("localhost:3000"),
				// "nginx.ingress.kubernetes.io/cors-allow-headers": jsii.String(
				// 	"x-user-agent, x-grpc-web, content-type, Sec-WebSocket-Extensions,Sec-WebSocket-Key, Sec-WebSocket-Protocol,Sec-WebSocket-Version",
				// ),
			},
			Name: jsii.String("ingress-api"),
		},
		Rules: &[]*cdk8splus.IngressRule{
			{
				Host:     jsii.String("example.api"),
				Path:     jsii.String("/markets(/|$)(.*)"),
				PathType: cdk8splus.HttpIngressPathType_PREFIX,
				Backend: cdk8splus.IngressBackend_FromService(props.MarketService, &cdk8splus.ServiceIngressBackendOptions{
					Port: jsii.Number(30002),
				}),
			},
		},
	})

	return chart
}
