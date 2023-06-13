package statefun

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	cdk8splus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2"
)

type StatefunProps struct {
	cdk8s.ChartProps
}

func SetStateFun(scope constructs.Construct, id string, props *StatefunProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	namespace := cdk8splus.NewNamespace(chart, jsii.String("statefun"), &cdk8splus.NamespaceProps{})
	_, k8sUrlService := SetStatefunFunctions(chart, "functions", &SetStatefunFunctionsProps{
		Namespace: &namespace,
	})

	appModuleConfigMap, appModuleVolume := SetAppModuleConfigMap(&chart, &SetAppModuleConfigMapProps{
		Namespace:    &namespace,
		PathFunction: k8sUrlService,
	})

	flinkConfigMap, flinkConfigVolume := SetFlinkConfigMap(&chart)

	SetJobManagerApp(&chart, SetJobManagerAppProps{
		Namespace:          &namespace,
		AppModuleConfigMap: &appModuleConfigMap,
		FlinkConfigMap:     &flinkConfigMap,
		AppModuleVolume:    &appModuleVolume,
		FlinkConfigVolume:  &flinkConfigVolume,
	})

	SetTaskJobManager(&chart, SetTaskJobManagerProps{
		Namespace:          &namespace,
		AppModuleConfigMap: &appModuleConfigMap,
		FlinkConfigMap:     &flinkConfigMap,
		AppModuleVolume:    &appModuleVolume,
		FlinkConfigVolume:  &flinkConfigVolume,
	})

	return chart
}
