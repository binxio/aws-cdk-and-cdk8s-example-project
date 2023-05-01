import cdk8s as cdk8s
import cdk8s_plus_25 as kplus

from constructs import Construct

class SampleChart(cdk8s.Chart):
    def __init__(self, scope: Construct, id: str):
        super().__init__(scope, id)

        nginx_deployment = kplus.Deployment(
            self, 'sampleDeployment',
            replicas=1,
            containers=[
                kplus.ContainerProps(
                    image='nginx:mainline-alpine',
                    port=80,
                    security_context=kplus.ContainerSecurityContextProps(
                        ensure_non_root=False,
                        read_only_root_filesystem=False
                    )
                )
            ],
            security_context=kplus.PodSecurityContextProps(
                ensure_non_root=False
            )
        )

        ingress = kplus.Ingress(
            self, 'alb'
        )

        nginx_deployment.expose_via_ingress(
            path='/',
            service_type=kplus.ServiceType.NODE_PORT,
            ingress=ingress
        )


