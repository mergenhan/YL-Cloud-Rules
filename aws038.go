package rules

const (
	AWSRuleID            = "038"
	AWSRuleDescription = "IAM Password policy should have expiry less than or equal to 90 days."
	AWSRelatedService = ["IAM"]
	AWSRule = "Critical"
	AWSRuleMitigation  = "Limit the password duration with an expiry in the policy
IAM account password policies should have a maximum age specified. 

The account password policy should be set to expire passwords after 90 days or less.
"
	AWSRuleItem = `
resource "aws_iam_account_password_policy"  {

	# max_password_age not set
	
}