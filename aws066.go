package rules

const AWSRuleID = "066"
const AWSRuleDescription = "EKS should have the encryption of secrets enabled"
const AWSRelatedService = ["EKS"]
const AWSRule = "Critical"

const AWSRuleItem = `
resource "aws_eks_cluster"  {
    name = "bad_example_cluster"

    role_arn = var.cluster_arn
    vpc_config {
        endpoint_public_access = false
    }
}
`