package rules

const (
	AWSRuleID            = "042"
	AWSRuleDescription = "IAM Password policy should have requirement for at least one lowercase character."
	AWSRelatedService = ["IAM"]
	AWSRule = "Critical"
	AWSRuleMitigation  = "Enforce longer, more complex passwords in the policy
IAM account password policies should ensure that passwords content including at least one lowercase character.
"
	AWSRuleItem = `
resource "aws_iam_account_password_policy"  {

	require_lowercase_characters not set

}
`