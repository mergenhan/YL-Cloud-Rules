package rules

const AWSRuleID = "069"
const AWSRuleDescription = "EKS Clusters should have the public access disabled"
const AWSRelatedService = ["EKS"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Don't enable public access to EKS Clusters
EKS clusters are available publicly by default, this should be explicitly disabled in the vpc_config of the EKS cluster resource.
"
const AWSRuleItem = `
resource "aws_eks_cluster" {
    // other config 

    name = "bad_example_cluster"
    role_arn = var.cluster_arn
    vpc_config {
		endpoint_public_access = true
		public_access_cidrs = ["0.0.0.0/0"]
    }
}
`