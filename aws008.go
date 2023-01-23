package rules

const AWSRuleID = "008"
const AWSRuleDescription = "An inline ingress security group rule allows traffic from /0."
const AWSRelatedService = ["SecurityHub", "API Gateway", "CloudFront"]
const AWSRule = "High"
const AWSRuleMitigation = "Set a more restrictive cidr range. Opening up ports to the public internet is generally to be avoided. You should restrict access to IP addresses or ranges that explicitly require it where possible.
"
const AWSRuleItem = `
resource "aws_security_group"  {
	ingress {
		cidr_blocks = ["0.0.0.0/0"]
	}
}