package rules


const AWSRuleID = "076"
const AWSRuleDescription = "S3 Access block should block public policy"
const AWSRelatedService = ["S3"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Prevent policies that allow public access being PUT
S3 bucket policy should have block public policy to prevent users from PUTing a policy that enable public access.
"
const AWSRuleItem = `
resource "aws_s3_bucket_public_access_block"  {
	bucket = aws_s3_bucket.example.id
}

resource "aws_s3_bucket_public_access_block" {
	bucket = aws_s3_bucket.example.id
  
	block_public_policy = false
}
`