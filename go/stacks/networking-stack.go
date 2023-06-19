package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	// https://docs.aws.amazon.com/sdk-for-go/api/aws/
	"github.com/aws/aws-sdk-go-v2/aws"

	"k8s-full-stack-cdk-go/helper"
)

type NetworkingNestedStackProps struct {
	awscdk.NestedStackProps
}

func NetworkingStack(scope constructs.Construct, id string, config helper.Conf, props *NetworkingNestedStackProps) (awscdk.NestedStack, awsec2.IVpc) {
	var nsprops awscdk.NestedStackProps
	if props != nil {
		nsprops = props.NestedStackProps
	}
	stack := awscdk.NewNestedStack(scope, &id, &nsprops)

	vpc := awsec2.NewVpc(stack, aws.String("k8s-sample-vpc"), 
		&awsec2.VpcProps{
			IpAddresses: awsec2.IpAddresses_Cidr(aws.String(config.CIDR)),
			SubnetConfiguration: &[]*awsec2.SubnetConfiguration{
				{
					CidrMask: aws.Float64(28),
					SubnetType: awsec2.SubnetType_PUBLIC,
					Name: aws.String("public"),
				},
				{
					CidrMask: aws.Float64(26),
					SubnetType: awsec2.SubnetType_PRIVATE_WITH_EGRESS,
					Name: aws.String("eks"),
				},
			},
		},
	)

	return stack, vpc
}