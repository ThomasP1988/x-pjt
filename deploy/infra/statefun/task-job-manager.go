package statefun

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	cdk8splus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2"
)

type SetTaskJobManagerProps struct {
	Namespace          *cdk8splus.Namespace
	FlinkConfigMap     *cdk8splus.ConfigMap
	AppModuleConfigMap *cdk8splus.ConfigMap
	AppModuleVolume    *cdk8splus.Volume
	FlinkConfigVolume  *cdk8splus.Volume
}

func SetTaskJobManager(chart *cdk8s.Chart, props SetTaskJobManagerProps) cdk8splus.Deployment {
	return cdk8splus.NewDeployment(*chart, jsii.String("task-job-manager"), &cdk8splus.DeploymentProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name:      jsii.String("statefun-worker"),
			Namespace: (*props.Namespace).Name(),
		},
		Replicas: jsii.Number(3),
		PodMetadata: &cdk8s.ApiObjectMetadata{
			Labels: &map[string]*string{
				"app":       jsii.String("statefun"),
				"component": jsii.String("worker"),
			},
		},
		Volumes: &[]cdk8splus.Volume{
			*props.AppModuleVolume,
			*props.FlinkConfigVolume,
		},
		Containers: &[]*cdk8splus.ContainerProps{
			{
				Name:            jsii.String("worker"),
				Image:           jsii.String("apache/flink-statefun:3.2.0"),
				ImagePullPolicy: cdk8splus.ImagePullPolicy_ALWAYS,
				EnvVariables: &map[string]cdk8splus.EnvValue{
					"ROLE":        cdk8splus.EnvValue_FromValue(jsii.String("worker")),
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
					Port:                jsii.Number(6122),
				}),
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
				Port: jsii.Number(6122),
			},
		},
	})
}
