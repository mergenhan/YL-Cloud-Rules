package rules

const AWSRuleID = "068"
const AWSRuleDescription = "EKS cluster should not have open CIDR range for public access"
const AWSRelatedService = ["EKS"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Don't enable public access to EKS Clusters
EKS Clusters have public access cidrs set to 0.0.0.0/0 by default which is wide open to the internet. This should be explicitly set to a more specific CIDR range
"
const AWSRuleItem = `
resource "aws_eks_cluster"  {
    // other config 

    name = "bad_example_cluster"
    role_arn = var.cluster_arn
    vpc_config {
        endpoint_public_access = true
    }
}
`