from aws_cdk import (
    # Duration,
    NestedStack,
    aws_ec2 as ec2,
)
from constructs import Construct
from helper import config

class NetworkingStack(NestedStack):

    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        conf = config.Config(self.node.try_get_context('environment'))

        self.vpc = ec2.Vpc(self, "k8s-sample-vpc",
            ip_addresses=ec2.IpAddresses.cidr(conf.get('cidr')),
            subnet_configuration=[
                ec2.SubnetConfiguration(
                    name = 'public',
                    subnet_type = ec2.SubnetType.PUBLIC,
                    cidr_mask = 28
                ),
                ec2.SubnetConfiguration(
                    name = 'eks',
                    subnet_type = ec2.SubnetType.PRIVATE_WITH_EGRESS,
                    cidr_mask = 26
                )
            ],
        )
