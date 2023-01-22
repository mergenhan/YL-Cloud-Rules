package rules

const AWSRuleID = 005
const AWSRuleDescription = "Plain HTTP is unencrypted and human-readable.
This means that if a malicious actor was to eavesdrop on your
connection, they would be able to see all of your data flowing back
and forth."
const AWSRuleImpact = "Critical"
const AWSRelatedService = ["EC2", "API Gateway"]
const AWSRuleMitigation = "Switch to HTTPS to benefit from TLS security
features. You should use HTTPS, which is HTTP over an encrypted (TLS)
connection, meaning eavesdroppers cannot read your traffic."
const AWSRuleItem =
resource "aws_network_protocol" { protocol = "HTTP"
}