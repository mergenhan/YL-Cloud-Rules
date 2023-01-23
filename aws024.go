package rules

const AWSRuleID = "024"
const AWSRuleDescription = "Kinesis stream is unencrypted."
const AWSRelatedService = ["Kinesis", "KBS", "CloudWatch"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enable in transit encryption 
Kinesis streams should be encrypted to ensure sensitive data is kept private. Additionally, non-default KMS keys should be used so granularity of access control can be ensured.
"
const AWSRuleItem = `
resource "aws_kinesis_stream" {
	encryption_type = "NONE"
}