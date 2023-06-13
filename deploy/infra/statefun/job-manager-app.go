package statefun

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	cdk8splus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2"
)

type SetJobManagerAppProps struct {
	Namespace          *cdk8splus.Namespace
	FlinkConfigMap     *cdk8splus.ConfigMap
	AppModuleConfigMap *cdk8splus.ConfigMap
	AppModuleVolume    *cdk8splus.Volume
	FlinkConfigVolume  *cdk8splus.Volume
}

func SetJobManagerApp(chart *cdk8s.Chart, props SetJobManagerAppProps) {

	deployment := cdk8splus.NewDeployment(*chart, jsii.String("job-manager-application"), &cdk8splus.DeploymentProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("statefun-jobmanager"),
		},
		Replicas: jsii.Number(1),
		Volumes: &[]cdk8splus.Volume{
			*props.AppModuleVolume,
			*props.FlinkConfigVolume,
		},
		PodMetadata: &cdk8s.ApiObjectMetadata{
			Labels: &map[string]*string{
				"app":       jsii.String("statefun"),
				"component": jsii.String("master"),
			},
		},
		Containers: &[]*cdk8splus.ContainerProps{
			{
				Name:            jsii.String("master"),
				Image:           jsii.String("apache/flink-statefun:3.2.0"),
				ImagePullPolicy: cdk8splus.ImagePullPolicy_ALWAYS,
				EnvVariables: &map[string]cdk8splus.EnvValue{
					"ROLE":        cdk8splus.EnvValue_FromValue(jsii.String("master")),
					"MASTER_HOST": cdk8splus.EnvValue_FromValue(jsii.String("statefun-jobmanager")),
				},
				Resources: &cdk8splus.ContainerResources{
					Memory: &cdk8splus.MemoryResources{
						Request: cdk8s.Size_Gibibytes(jsii.Number(1.5)),
					},
				},
				Liveness: cdk8splus.Probe_FromTcpSocket(&cdk8splus.TcpSocketProbeOptions{
					PeriodSeconds:       cdk8s.Duration_Seconds(jsii.Number(60)),
					InitialDelaySeconds: cdk8s.Duration_Seconds(jsii.Number(30)),
					Port:                jsii.Number(6123),
				}),
				Port: jsii.Number(6123),
				VolumeMounts: &[]*cdk8splus.VolumeMount{
					{
						Path:   jsii.String("/opt/flink/conf"),
						Volume: *props.FlinkConfigVolume,
					},
					{
						Path:   jsii.String("/opt/statefun/modules/application-module"),
						Volume: *props.AppModuleVolume,
					},
				},
			},
		},
	})

	deployment.ExposeViaService(&cdk8splus.DeploymentExposeViaServiceOptions{
		Name:        jsii.String("flink-jobmanager"),
		ServiceType: cdk8splus.ServiceType_CLUSTER_IP,
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

	deployment.ExposeViaService(&cdk8splus.DeploymentExposeViaServiceOptions{
		Name:        jsii.String("statefun-jobmanager-rest"),
		ServiceType: cdk8splus.ServiceType_NODE_PORT,
		Ports: &[]*cdk8splus.ServicePort{
			{
				Name:       jsii.String("rest"),
				Port:       jsii.Number(8081),
				TargetPort: jsii.Number(8081),
				NodePort:   jsii.Number(30081),
			},
		},
	})

}
