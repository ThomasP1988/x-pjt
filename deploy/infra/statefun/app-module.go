package statefun

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	cdk8splus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2"
)

type SetAppModuleConfigMapProps struct {
	Namespace    *cdk8splus.Namespace
	PathFunction string
}

func SetAppModuleConfigMap(chart *cdk8s.Chart, props *SetAppModuleConfigMapProps) (cdk8splus.ConfigMap, cdk8splus.Volume) {

	configmap := cdk8splus.NewConfigMap(*chart, jsii.String("statefun-configmap"), &cdk8splus.ConfigMapProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name:      jsii.String("application-module"),
			Namespace: (*props.Namespace).Name(),
			Labels: &map[string]*string{
				"app": jsii.String("statefun"),
			},
		},
		Data: &map[string]*string{
			"module.yaml": jsii.String(`kind: io.statefun.endpoints.v2/http
spec:
    functions: com.nftm/*
    urlPathTemplate: ` + props.PathFunction + `/{function.name}
`),
		},
	})

	volume := cdk8splus.Volume_FromConfigMap(*chart, jsii.String("job-manager-app-module-config"), configmap, &cdk8splus.ConfigMapVolumeOptions{
		Name: jsii.String("application-module"),
		Items: &map[string]*cdk8splus.PathMapping{
			"module.yaml": {
				Path: jsii.String("module.yaml"),
			},
		},
	})

	return configmap, volume
}
