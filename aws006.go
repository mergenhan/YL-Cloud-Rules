package rules

const AWSRuleID = "006"
const AWSRuleDescription = "An ingress security group rule allows traffic from /0."
const AWSRelatedService = ["EC2", "API Gateway", "CloudFront"]
const AWSRuleImpact = "High"
const AWSRuleMitigation = "Set a more restrictive cidr range"

const AWSRuleItem = `
resource "aws_security_group_rule" {
	type = "ingress"
	cidr_blocks = ["0.0.0.0/0"]
}