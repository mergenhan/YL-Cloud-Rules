package rules

const AWSRuleID = 002
const AWSRuleDescription = "An security group rule allows traffic from
anywhere. Opening up ports to the public internet is generally to be
avoided. You should restrict access to IP addresses or ranges that
explicitly require it where possible."
const AWSRelatedService = ["EC2", "API Gateway", "CloudFront"]
const AWSRuleImpact = "Critical"
const AWSRuleMitigation = "Set a more restrictive cidr range"
const AWSRuleItem =
resource "$aws_security_group_rule" {
cidr_blocks = ["0.0.0.0/0"]