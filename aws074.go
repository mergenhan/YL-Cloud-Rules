package rules

const AWSRuleID = "074"
const AWSRuleDescription = "S3 Access block should block public ACL"
const AWSRelatedService = ["S3"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enable blocking any PUT calls with a public ACL specified
S3 buckets should block public ACLs on buckets and any objects they contain. By blocking, PUTs with fail if the object has any public ACL a.
"
const AWSRuleItem = `
resource "aws_s3_bucket_public_access_block"  {
	bucket = aws_s3_bucket.example.id
}

resource "aws_s3_bucket_public_access_block" {
	bucket = aws_s3_bucket.example.id
  
	block_public_acls = false
}
`