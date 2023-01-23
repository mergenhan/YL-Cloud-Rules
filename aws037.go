package rules

const (
	AWSRuleID            = "037"
	AWSRuleDescription = "IAM Password policy should prevent password reuse."
	AWSRelatedService = ["IAM"]
	AWSRule = "Critical"
	AWSRuleMitigation  = "Prevent password reuse in the policy 
IAM account password policies should prevent the reuse of passwords. 

The account password policy should be set to prevent using any of the last five used passwords.
"
	AWSRuleItem = `
resource "aws_iam_account_password_policy"  {

	password_reuse_prevention = 1
	
}
`