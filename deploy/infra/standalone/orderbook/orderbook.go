package orderbook

import (
	"NFTM/shared/components/market"
	"fmt"
	"infra/aws"
	"log"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	cdk8splus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2"
)

type MyChartProps struct {
	cdk8s.ChartProps
	Market market.MarketConfig
}

func SetOrderbook(scope constructs.Construct, id string, props *MyChartProps) (cdk8splus.Service, cdk8s.Chart) {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// TODO: use iam when we will be having EKS
	fmt.Printf("props.config: %v\n", props.Market)
	credentials, err := aws.GetCredentials()

	if err != nil {
		log.Fatalf(err.Error())
	}

	deploymentName := jsii.String("orderbook-" + props.Market.Pair.StringLowercase())

	deployment := cdk8splus.NewDeployment(chart, jsii.String("orderbook-deployement-"+props.Market.Pair.StringLowercase()), &cdk8splus.DeploymentProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: deploymentName,
			Labels: &map[string]*string{
				"app": deploymentName,
			},
		},
		Replicas: jsii.Number(1),
		Containers: &[]*cdk8splus.ContainerProps{{
			Image:           jsii.String("orderbook"),
			ImagePullPolicy: cdk8splus.ImagePullPolicy_NEVER,
			// TODO: remove when using IAM
			EnvVariables: &map[string]cdk8splus.EnvValue{
				"AWS_ACCESS_KEY_ID":     cdk8splus.EnvValue_FromValue(jsii.String(credentials.AccessKeyID)),
				"AWS_SECRET_ACCESS_KEY": cdk8splus.EnvValue_FromValue(jsii.String(credentials.SecretAccessKey)),
				"MARKET_SYMBOL":         cdk8splus.EnvValue_FromValue(jsii.String(props.Market.Pair.Symbol())),
			},
		}},
		RestartPolicy: cdk8splus.RestartPolicy_ALWAYS,
	})

	service := deployment.ExposeViaService(&cdk8splus.DeploymentExposeViaServiceOptions{
		Name:        jsii.String("orderbook-grpc-" + props.Market.Pair.StringLowercase()),
		ServiceType: cdk8splus.ServiceType_CLUSTER_IP,
		Ports: &[]*cdk8splus.ServicePort{
			{
				Name: jsii.String(props.Market.DNS),
				// NodePort:   jsii.Number(30002),
				Protocol:   cdk8splus.Protocol_TCP,
				TargetPort: jsii.Number(50052),
				Port:       jsii.Number(float64(props.Market.Port)),
			},
		},
	})

	return service, chart
}
