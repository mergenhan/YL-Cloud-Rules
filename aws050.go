package rules

const AWSRuleID = "050"
const AWSRuleDescription = "An ingress Network ACL rule allows ALL ports from /0."
const AWSRelatedService = ["API Gateway", "EC2"]
const AWSRule = "Critical"

const AWSRuleItem = `
resource "aws_network_acl_rule" {
  egress         = false
  protocol       = "all"
  rule_action    = "allow"
  cidr_block     = "0.0.0.0/0"
}
`