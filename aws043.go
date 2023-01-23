package rules

const (
	AWSRuleID            = "043"
	AWSRuleDescription = "IAM Password policy should have requirement for at least one uppercase character."
	AWSRelatedService = ["IAM"]
	AWSRule = "Critical"
	AWSRuleMitigation  = "Enforce longer, more complex passwords in the policy
IAM account password policies should ensure that passwords content including at least one uppercase character.
"
	AWSRuleItem = `
resource "aws_iam_account_password_policy"  {
	# ...
	# require_uppercase_characters not set
	# ...
}
`