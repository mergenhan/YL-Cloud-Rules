package rules


const AWSRuleID = "049"
const AWSRuleDescription = "An ingress Network ACL rule allows specific ports from /0."
const AWSRelatedService = ["EC2","API Gateway", "CloudWatch"]
const AWSRule = "Medium"

const AWSRuleItem = `
resource "aws_network_acl_rule"  {
  egress         = false
  protocol       = "tcp"
  from_port      = 22
  to_port        = 22
  rule_action    = "allow"
  cidr_block     = "0.0.0.0/0"
}