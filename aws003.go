package rules

const AWSRuleID = "003"
const AWSRuleDescription = "AWS Classic resource usage."
const AWSRuleImpact = "High"
const AWSRuleMitigation = "Switch to VPC resources"

const AWSRuleItem = `
resource "aws_db_security_group" "bad_example" {
  # ...
}
