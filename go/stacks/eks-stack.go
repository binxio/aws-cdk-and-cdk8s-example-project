package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	// https://docs.aws.amazon.com/sdk-for-go/api/aws/
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cdklabs/awscdk-kubectl-go/kubectlv25/v2"

	"k8s-full-stack-cdk-go/helper"
)

type EksNestedStackProps struct {
	awscdk.NestedStackProps
	Vpc awsec2.IVpc
}

func EKSStack(scope constructs.Construct, id string, config helper.Conf, props *EksNestedStackProps) (awscdk.NestedStack, awseks.ICluster) {
	var nsprops awscdk.NestedStackProps
	if props != nil {
		nsprops = props.NestedStackProps
	}
	stack := awscdk.NewNestedStack(scope, &id, &nsprops)

	var cluster awseks.Cluster

	cluster = awseks.NewCluster(stack, aws.String("k8s-sample-cluster"), &awseks.ClusterProps{
		Version: awseks.KubernetesVersion_V1_25(),
		KubectlLayer: kubectlv25.NewKubectlV25Layer(stack, aws.String("kubectl")),
		AlbController: &awseks.AlbControllerOptions{
			Version: awseks.AlbControllerVersion_V2_4_1(),
		},
		DefaultCapacity: aws.Float64(0),
		Vpc: props.Vpc,
		VpcSubnets: &[]*awsec2.SubnetSelection{
			{
				SubnetGroupName: aws.String("eks"),
			},
		},
	})

	cluster.AddNodegroupCapacity(aws.String("eks-nodegroup"), &awseks.NodegroupOptions{
		InstanceTypes: &[]awsec2.InstanceType{
			awsec2.NewInstanceType(aws.String("t3.large")),
		},
	})

	adminUser := awsiam.User_FromUserArn(stack, aws.String("imported_user"), aws.String("arn:aws:iam::<account_id>:user/<iam_user>"))
	cluster.AwsAuth().AddUserMapping(adminUser, &awseks.AwsAuthMapping{
		Groups: &[]*string{
			aws.String("system:masters"),
		},
	})
	

	return stack, cluster
}
