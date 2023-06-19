from aws_cdk import (
    # Duration,
    NestedStack,
    aws_eks as eks,
    aws_ec2 as ec2,
    aws_iam as iam
)
from constructs import Construct
from helper import config
from aws_cdk.lambda_layer_kubectl_v25 import KubectlV25Layer

class EKSStack(NestedStack):

    def __init__(self, scope: Construct, construct_id: str, vpc: ec2.IVpc, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        conf = config.Config(self.node.try_get_context('environment'))

        self.cluster = eks.Cluster(
            self, 'k8s-sample-cluster',
            version=eks.KubernetesVersion.V1_25,
            kubectl_layer=KubectlV25Layer(self, "kubectl"),
            alb_controller=eks.AlbControllerOptions(
                version=eks.AlbControllerVersion.V2_4_1
            ),
            default_capacity=0,
            vpc=vpc,
            vpc_subnets=[
                ec2.SubnetSelection(
                    subnet_group_name='eks'
                )
            ]
        )

        self.cluster.add_nodegroup_capacity(
            'eks-nodegroup',
            instance_types=[ec2.InstanceType('t3.large')]
        )

        # importing existing iam user
        admin_user = iam.User.from_user_arn(
            self, 'imported_user',
            user_arn='arn:aws:iam::<account_id>:user/<iam_user>' # change me
        )
        self.cluster.aws_auth.add_user_mapping(admin_user, groups=["system:masters"])



