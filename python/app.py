#!/usr/bin/env python3
import os

import aws_cdk as cdk
import cdk8s as cdk8s
from helper import config

from k8s_full_stack_cdk.k8s_full_stack_cdk_stack import K8SFullStackCdkStack
from k8s_full_stack_cdk.networking_stack import NetworkingStack
from k8s_full_stack_cdk.eks_stack import EKSStack
from k8s_full_stack_cdk.eks_application_stack import EKSApplicationStack
from k8s_full_stack_charts.sample_chart import SampleChart


app = cdk.App()
cdk8s_app = cdk8s.App()

my_chart = SampleChart(
    cdk8s_app, "SampleChart"
)

conf = config.Config(app.node.try_get_context('environment'))
k8s_sample = K8SFullStackCdkStack(app, "K8SFullStackCdkStack",
    # If you don't specify 'env', this stack will be environment-agnostic.
    # Account/Region-dependent features and context lookups will not work,
    # but a single synthesized template can be deployed anywhere.

    # Uncomment the next line to specialize this stack for the AWS Account
    # and Region that are implied by the current CLI configuration.

    #env=cdk.Environment(account=os.getenv('CDK_DEFAULT_ACCOUNT'), region=os.getenv('CDK_DEFAULT_REGION')),

    # Uncomment the next line if you know exactly what Account and Region you
    # want to deploy the stack to. */

    env=cdk.Environment(account=conf.get('account_id'), region=conf.get('region')),

    # For more information, see https://docs.aws.amazon.com/cdk/latest/guide/environments.html
    )

networking = NetworkingStack(
    k8s_sample, 'NetworkingStack'
)

eks_cluster = EKSStack(
    k8s_sample, 'EKSStack',
    vpc = networking.vpc
)

eks_application = EKSApplicationStack(
    k8s_sample, 'EKSApplicationStack',
    cluster = eks_cluster.cluster,
    chart = my_chart
)


app.synth()
