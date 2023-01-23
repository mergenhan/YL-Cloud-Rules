package rules

const AWSRuleID = "004"
const AWSRuleDescription = "Use of plain HTTP."
const AWSRelatedService = ["EC2", "API Gateway"]
const AWSRuleImpact = "High"
const AWSRuleMitigation = "Switch to HTTPS to benefit from TLS security features"

const AWSRuleItem = `
resource "aws_alb_listener" {
	protocol = "HTTP"
}
`