package rules

const (
	AWSRuleID            = "040"
	AWSRuleDescription = "IAM Password policy should have requirement for at least one symbol in the password."
	AWSRelatedService = ["IAM"]
	AWSRule = "Critical"
	AWSRuleMitigation  = "Enforce longer, more complex passwords in the policy 
IAM account password policies should ensure that passwords content including a symbol.
"
	AWSRuleItem = `
resource "aws_iam_account_password_policy" {

	# require_symbols not set

}