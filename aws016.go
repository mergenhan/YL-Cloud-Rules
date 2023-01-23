package rules


const AWSRuleID = "016"
const AWSRuleDescription = "Unencrypted SNS topic."
const AWSRelatedService = ["SNS"]
const AWSRule = "High"
const AWSRuleMitigation = "Turn on SNS Topic encryption 
Queues should be encrypted with customer managed KMS keys and not default AWS managed keys, in order to allow granular control over access to specific queues.
"
const AWSRuleItem = `
resource "aws_sns_topic"  {
	# no key id specified
}