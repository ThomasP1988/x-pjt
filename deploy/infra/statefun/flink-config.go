package statefun

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	cdk8splus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2"
)

func SetFlinkConfigMap(chart *cdk8s.Chart) (cdk8splus.ConfigMap, cdk8splus.Volume) {
	configmap := cdk8splus.NewConfigMap(*chart, jsii.String("flink-configmap"), &cdk8splus.ConfigMapProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("flink-config"),
			Labels: &map[string]*string{
				"app": jsii.String("statefun"),
			},
		},
		Data: &map[string]*string{
			"flink-conf.yaml": jsii.String(`jobmanager.rpc.address: statefun-master
taskmanager.numberOfTaskSlots: 1
blob.server.port: 6124
jobmanager.rpc.port: 6123
taskmanager.rpc.port: 6122
classloader.parent-first-patterns.additional: org.apache.flink.statefun;org.apache.kafka;com.google.protobuf
jobmanager.memory.process.size: 1g
taskmanager.memory.process.size: 1g
`),
		},
	})

	volume := cdk8splus.Volume_FromConfigMap(*chart, jsii.String("job-manager-app-flink-config"), configmap, &cdk8splus.ConfigMapVolumeOptions{
		Name: jsii.String("flink-config"),
		Items: &map[string]*cdk8splus.PathMapping{
			"flink-conf.yaml": {
				Path: jsii.String("flink-conf.yaml"),
			},
			"log4j-console.properties": {
				Path: jsii.String("log4j-console.properties"),
			},
		},
	})

	return configmap, volume
}
