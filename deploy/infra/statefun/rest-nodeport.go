package statefun

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	cdk8splus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2"
)

func SetJobManagerRestService(chart *cdk8s.Chart, selector cdk8splus.IPodSelector) cdk8splus.Service {

	service := cdk8splus.NewService(*chart, jsii.String("jobmanager-rest-nodeport"), &cdk8splus.ServiceProps{
		Type: cdk8splus.ServiceType_NODE_PORT,
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("statefun-jobmanager-rest"),
		},
		Selector: selector,
		Ports: &[]*cdk8splus.ServicePort{
			{
				Name:       jsii.String("rest"),
				Port:       jsii.Number(8081),
				TargetPort: jsii.Number(8081),
				NodePort:   jsii.Number(30081),
			},
		},
	})

	return service
}
