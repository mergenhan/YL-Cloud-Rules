package rules

const AWSRuleID = "017"
const AWSRuleDescription = "Unencrypted S3 bucket."
const AWSRelatedService = ["S3"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Configure bucket encryption 
S3 Buckets should be encrypted with customer managed KMS keys and not default AWS managed keys, in order to allow granular control over access to specific buckets.
"
const AWSRuleItem = `
resource "aws_s3_bucket"  {
  bucket = "mybucket"
}
`