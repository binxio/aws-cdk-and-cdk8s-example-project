package charts

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus25/v2"
)

type SampleChartProps struct {
	cdk8s.ChartProps
}

func SampleChart(scope constructs.Construct, id *string, props *SampleChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, id, &cprops)

	nginxDeployment := cdk8splus25.NewDeployment(chart, aws.String("sampleDeployment"), &cdk8splus25.DeploymentProps{
		Replicas: aws.Float64(1),
		Containers: &[]*cdk8splus25.ContainerProps{
			{
				Image: aws.String("nginx:mainline-alpine"),
				Port: aws.Float64(80),
				SecurityContext: &cdk8splus25.ContainerSecurityContextProps{
					EnsureNonRoot: aws.Bool(false),
					ReadOnlyRootFilesystem: aws.Bool(false),
				},
			},
		},
		SecurityContext: &cdk8splus25.PodSecurityContextProps{
			EnsureNonRoot: aws.Bool(false),
		},
	})

	ingress := cdk8splus25.NewIngress(chart, aws.String("alb"), nil)

	nginxDeployment.ExposeViaIngress(aws.String("/"), &cdk8splus25.ExposeDeploymentViaIngressOptions{
		ServiceType: cdk8splus25.ServiceType_NODE_PORT,
		Ingress: ingress,
	})


	return chart
}