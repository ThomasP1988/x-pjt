package statefun

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	cdk8splus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2"
)

func SetJobManagerService(chart *cdk8s.Chart, selector cdk8splus.IPodSelector) cdk8splus.Service {
	return cdk8splus.NewService(*chart, jsii.String("jobmanager-service"), &cdk8splus.ServiceProps{
		Type: cdk8splus.ServiceType_CLUSTER_IP,
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("flink-jobmanager"),
		},
		Selector: selector,
		Ports: &[]*cdk8splus.ServicePort{
			{
				Name: jsii.String("rpc"),
				Port: jsii.Number(6123),
			},
			{
				Name: jsii.String("blob-server"),
				Port: jsii.Number(6124),
			},
			{
				Name: jsii.String("webui"),
				Port: jsii.Number(8081),
			},
		},
	})
}
