package grpc

import (
	"NFTM/shared/config"
	"infra/aws"
	"log"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	cdk8splus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2"
)

type MyChartProps struct {
	cdk8s.ChartProps
}

func SetAPIMarket(scope constructs.Construct, id string, props *MyChartProps) (cdk8splus.Service, cdk8s.Chart) {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	deploymentName := jsii.String("apis-grpc")

	// TODO: use iam when we will be having EKS

	credentials, err := aws.GetCredentials()

	if err != nil {
		log.Fatalf(err.Error())
	}

	deployment := cdk8splus.NewDeployment(chart, jsii.String("market-api-deployment"), &cdk8splus.DeploymentProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: deploymentName,
			Labels: &map[string]*string{
				"app": deploymentName,
			},
		},
		Replicas: jsii.Number(1),
		Containers: &[]*cdk8splus.ContainerProps{{
			Image:           jsii.String("apis/grpc"),
			ImagePullPolicy: cdk8splus.ImagePullPolicy_NEVER,
			// TODO: remove when using IAM
			EnvVariables: &map[string]cdk8splus.EnvValue{
				"AWS_ACCESS_KEY_ID":     cdk8splus.EnvValue_FromValue(jsii.String(credentials.AccessKeyID)),
				"AWS_SECRET_ACCESS_KEY": cdk8splus.EnvValue_FromValue(jsii.String(credentials.SecretAccessKey)),
			},
			Port: jsii.Number(50052), //TODO add more port when CDK8s allows it
		}},
		RestartPolicy: cdk8splus.RestartPolicy_ALWAYS,
	})

	service := deployment.ExposeViaService(&cdk8splus.DeploymentExposeViaServiceOptions{
		ServiceType: cdk8splus.ServiceType_CLUSTER_IP,
		Name:        jsii.String("market-grpc-web-http-port"),
		Ports: &[]*cdk8splus.ServicePort{
			{
				Name: jsii.String("market-grpc-port"),
				// NodePort:   jsii.Number(30001),
				TargetPort: jsii.Number(50051),
				Protocol:   cdk8splus.Protocol_TCP,
				Port:       jsii.Number(float64(config.Conf.Apis[config.MarketGRPC].Port)),
			},
			{
				Name: jsii.String("market-grpc-web-http-port"),
				// NodePort:   jsii.Number(30002),
				TargetPort: jsii.Number(50052),
				Protocol:   cdk8splus.Protocol_TCP,
				Port:       jsii.Number(float64(config.Conf.Apis[config.MarketGRPCHTTP].Port)),
			},
		},
	})

	return service, chart
}
