package rules

const AWSRuleID = "073"
const AWSRuleDescription = "S3 Access Block should Ignore Public Acl"
const AWSRelatedService = ["S3"]
const AWSRule = "High"
const AWSRuleMitigation = "Enable ignoring the application of public ACLs in PUT calls
S3 buckets should ignore public ACLs on buckets and any objects they contain. By ignoring rather than blocking, PUT calls with public ACLs will still be applied but the ACL will be ignored.
"
const AWSRuleItem = `
resource "aws_s3_bucket_public_access_block"  {
	bucket = aws_s3_bucket.example.id
}

resource "aws_s3_bucket_public_access_block" "" {
	bucket = aws_s3_bucket.example.id
  
	ignore_public_acls = false
}
`