package rules

const AWSRuleID = "001"
const AWSRuleDescription = "S3 Bucket has an ACL defined which allows public access."
const AWSRelatedService = ["S3"]
const AWSRuleImpact = "Critical"
const AWSRuleMitigation = "Apply a more restrictive bucket ACL"

const AWSRuleItem = `
resource "aws_s3_bucket"  {
	acl = "public-read"
}
`