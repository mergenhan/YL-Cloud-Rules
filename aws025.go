package rules

const AWSRuleID = "025"
const AWSRuleDescription = "API Gateway domain name uses outdated SSL/TLS protocols."
const AWSRelatedService = ["Kinesis", "API Gateway", "CloudWatch"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Use the most modern TLS/SSL policies available You should not use outdated/insecure TLS versions for encryption. You should be using TLS v1.2+.
"
const AWSRuleItem = `
resource "aws_api_gateway_domain_name"  {
	security_policy = "TLS_1_0"
}
`