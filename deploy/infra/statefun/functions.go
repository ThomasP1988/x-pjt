package statefun

import (
	"infra/aws"
	"log"
	"strconv"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	cdk8splus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2"
)

type SetStatefunFunctionsProps struct {
	Namespace *cdk8splus.Namespace
}

func SetStatefunFunctions(chart constructs.Construct, id string, props *SetStatefunFunctionsProps) (cdk8splus.Service, string) {

	deploymentName := jsii.String("statefun-functions")
	httpPort := 8000
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
			Namespace: (*props.Namespace).Name(),
		},
		Replicas: jsii.Number(1),
		Containers: &[]*cdk8splus.ContainerProps{{
			Image:           jsii.String("statefun/functions"),
			ImagePullPolicy: cdk8splus.ImagePullPolicy_NEVER,
			// TODO: remove when using IAM
			EnvVariables: &map[string]cdk8splus.EnvValue{
				"AWS_ACCESS_KEY_ID":     cdk8splus.EnvValue_FromValue(jsii.String(credentials.AccessKeyID)),
				"AWS_SECRET_ACCESS_KEY": cdk8splus.EnvValue_FromValue(jsii.String(credentials.SecretAccessKey)),
			},
			Port: jsii.Number(float64(httpPort)), //TODO add more port when CDK8s allows it
		}},
		RestartPolicy: cdk8splus.RestartPolicy_ALWAYS,
	})

	service := deployment.ExposeViaService(&cdk8splus.DeploymentExposeViaServiceOptions{
		ServiceType: cdk8splus.ServiceType_CLUSTER_IP,
		Name:        jsii.String("statefun-function-http"),
		Ports: &[]*cdk8splus.ServicePort{
			{
				Name: jsii.String("statefun-function-http-port"),
				// NodePort:   jsii.Number(30002),
				TargetPort: jsii.Number(float64(httpPort)),
				Protocol:   cdk8splus.Protocol_TCP,
				Port:       jsii.Number(float64(httpPort)),
			},
		},
	})

	k8sUrlService := *service.Name() + "." + *(*props.Namespace).Name() + ".svc.cluster.local:" + strconv.Itoa(httpPort)

	return service, k8sUrlService
}
