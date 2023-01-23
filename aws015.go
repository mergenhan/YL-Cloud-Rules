package rules


const AWSRuleID = "015"
const AWSRuleDescription = "Unencrypted SQS queue."
const AWSRelatedService = ["SQS"]
const AWSRule = "Medium"
const AWSRuleMitigation = "Turn on SQS Queue encryption 
Queues should be encrypted with customer managed KMS keys and not default AWS managed keys, in order to allow granular control over access to specific queues.
"
const AWSRuleItem = `
resource "aws_sqs_queue"  {
	# no key specified
}