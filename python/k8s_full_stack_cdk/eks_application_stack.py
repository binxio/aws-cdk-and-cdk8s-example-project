from aws_cdk import (
    # Duration,
    NestedStack,
    aws_eks as eks
)
import cdk8s as cdk8s
from constructs import Construct
from helper import config

class EKSApplicationStack(NestedStack):

    def __init__(self, scope: Construct, construct_id: str, cluster: eks.ICluster, chart: cdk8s.Chart, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        conf = config.Config(self.node.try_get_context('environment'))

        cluster.add_cdk8s_chart(
            'sample-chart',
            chart,
            ingress_alb=True,
            ingress_alb_scheme=eks.AlbScheme.INTERNET_FACING,
            prune=True
        )

