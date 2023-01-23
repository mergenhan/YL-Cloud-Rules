package rules

const (
	AWSRuleID            = "041"
	AWSRuleDescription = "IAM Password policy should have requirement for at least one number in the password."
	AWSRelatedService = ["IAM"]
	AWSRule = "Critical"
	AWSRuleMitigation  = "Enforce longer, more complex passwords in the policy  
IAM account password policies should ensure that passwords content including at least one number.
"
	AWSRuleItem = `
resource "aws_iam_account_password_policy" {
	# require_numbers not set
	
}