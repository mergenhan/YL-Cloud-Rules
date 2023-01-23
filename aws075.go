package rules


const AWSRuleID = "075"
const AWSRuleDescription = "S3 Access block should restrict public bucket to limit access"
const AWSRelatedService = ["S3"]
const AWSRule = "High"
const AWSRuleMitigation = "Limit the access to public buckets to only the owner or AWS Services (eg; CloudFront).
S3 buckets should restrict public policies for the bucket. By enabling, the restrict_public_buckets, only the bucket owner and AWS Services can access if it has a public policy.
"
const AWSRuleItem = `
resource "aws_s3_bucket_public_access_block" {
	bucket = aws_s3_bucket.example.id
}

resource "aws_s3_bucket_public_access_block"  {
	bucket = aws_s3_bucket.example.id
  
	restrict_public_buckets = false
}
`