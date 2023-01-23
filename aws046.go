package rules

const AWSRuleID = "046"
const AWSRuleDescription = "AWS IAM policy document has wildcard action statement."
const AWSRelatedService = ["IAM"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Keep policy scope to the minimum that is required to be effective
IAM profiles should be configured with the specific, minimum set of permissions required.
"
const AWSRuleItem = `
data "aws_iam_policy_document"  {
	statement {
		sid = "1"

        actions = [
      		"*"
    	]
	}
}
`