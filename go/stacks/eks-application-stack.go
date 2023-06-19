package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/constructs-go/constructs/v10"
	// https://docs.aws.amazon.com/sdk-for-go/api/aws/
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"

	"k8s-full-stack-cdk-go/helper"
	"k8s-full-stack-cdk-go/charts"
)

type EksApplicationNestedStackProps struct {
	awscdk.NestedStackProps
	Cluster awseks.ICluster
}

func EKSApplicationStack(scope constructs.Construct, id string, config helper.Conf, props *EksApplicationNestedStackProps) awscdk.NestedStack {
	var nsprops awscdk.NestedStackProps
	if props != nil {
		nsprops = props.NestedStackProps
	}
	stack := awscdk.NewNestedStack(scope, &id, &nsprops)
	cluster := props.Cluster

	chart := charts.SampleChart(cdk8s.NewApp(nil), aws.String("SampleChart"), nil)
	cluster.AddCdk8sChart(aws.String("sample-chart"), chart, &awseks.KubernetesManifestOptions{
		IngressAlb: aws.Bool(true),
		IngressAlbScheme: awseks.AlbScheme_INTERNET_FACING,
		Prune: aws.Bool(true),
	})

	return stack
}
