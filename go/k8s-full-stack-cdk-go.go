package main

import (
	"io/ioutil"
	"k8s-full-stack-cdk-go/helper"
	"k8s-full-stack-cdk-go/stacks"
	"log"
	"reflect"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	// https://docs.aws.amazon.com/sdk-for-go/api/aws/
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/jsii-runtime-go"
)

// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"

type K8SFullStackCdkGoStackProps struct {
	awscdk.StackProps
}

func NewK8SFullStackCdkGoStack(scope constructs.Construct, id string, config helper.Conf, props *K8SFullStackCdkGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, aws.String("K8SFullStackCdkGoQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(aws.Number(300)),
	// })

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	// get environment from context
	envName := app.Node().GetContext(aws.String("environment"))
	envNameString := reflect.ValueOf(envName).String()

	stack := NewK8SFullStackCdkGoStack(app, "K8SFullStackCdkGoStack", ReadEnvFile(envNameString), &K8SFullStackCdkGoStackProps{
		awscdk.StackProps{
			Env: env(reflect.ValueOf(envName).String()),
		},
	})

	_, nsvpc := stacks.NetworkingStack(stack, "NetworkingStack", ReadEnvFile(envNameString), nil)
	_, cluster := stacks.EKSStack(stack, "EKSStack", ReadEnvFile(envNameString), &stacks.EksNestedStackProps{
		Vpc: nsvpc,
	})
	stacks.EKSApplicationStack(stack, "EKSApplicationStack", ReadEnvFile(envNameString), &stacks.EksApplicationNestedStackProps{
		Cluster: cluster,
	})

	app.Synth(nil)
}

func ReadEnvFile(environment string) helper.Conf {
	// retrieve configuration
	conf, err := ioutil.ReadFile("config/" + environment + ".yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	return helper.Config(conf)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env(environment string) *awscdk.Environment {

	// retrieve configuration
	config := ReadEnvFile(environment)

	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	// return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: aws.String(config.AccountID),
		Region:  aws.String(config.Region),
	}

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: aws.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  aws.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
