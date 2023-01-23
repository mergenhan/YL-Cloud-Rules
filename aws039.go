package rules

const (
	AWSRuleID            = "039"
	AWSRuleDescription = "IAM Password policy should have minimum password length of 14 or more characters."
	AWSRelatedService = ["IAM"]
	AWSRule = "Critical"
	AWSRuleMitigation  = "Enforce longer, more complex passwords in the policy 
IAM account password policies should ensure that passwords have a minimum length. 

The account password policy should be set to enforce minimum password length of at least 14 characters.
"
	AWSRuleItem = `
resource "aws_iam_account_password_policy" "bad_example" {

	# minimum_password_length not set

}