package rules

const AWSRuleID = "004"
const AWSRuleDescription = "Use of plain HTTP."
const AWSRuleImpact = "High"
const AWSRuleMitigation = "Switch to HTTPS to benefit from TLS security features"

const AWSRuleItem = `
resource "aws_alb_listener" "bad_example" {
	protocol = "HTTP"
}
`