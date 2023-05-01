import aws_cdk as core
import aws_cdk.assertions as assertions

from k8s_full_stack_cdk.k8s_full_stack_cdk_stack import K8SFullStackCdkStack

# example tests. To run these tests, uncomment this file along with the example
# resource in k8s_full_stack_cdk/k8s_full_stack_cdk_stack.py
def test_sqs_queue_created():
    app = core.App()
    stack = K8SFullStackCdkStack(app, "k8s-full-stack-cdk")
    template = assertions.Template.from_stack(stack)

#     template.has_resource_properties("AWS::SQS::Queue", {
#         "VisibilityTimeout": 300
#     })
