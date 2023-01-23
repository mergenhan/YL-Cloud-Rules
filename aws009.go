package rules

const AWSRuleID = "009"
const AWSRuleDescription = "An inline egress security group rule allows traffic to /0."
const AWSRelatedService = ["EC2", "API Gateway", "CloudFront"]
const AWSRule = "High"
const AWSRuleMitigation =  "Set a more restrictive cidr range 
Opening up ports to the public internet is generally to be avoided. You should restrict access to IP addresses or ranges that explicitly require it where possible.
"
const AWSRuleItem = `
resource "aws_security_group"   {
	egress {
		cidr_blocks = ["0.0.0.0/0"]
	}
}