package rules

const AWSRuleID = "019"
const AWSRuleDescription = "A KMS key is not configured to auto-rotate."
const AWSRelatedService = ["KMS"]
const AWSRule = "Critical"
const AWSRuleMitigation ="Configure KMS key to auto rotate 
You should configure your KMS keys to auto rotate to maintain security and defend against compromise.
"
const AWSRuleItem = `
resource "aws_kms_key"  {
	enable_key_rotation = false
}
`