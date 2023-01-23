package rules

const AWSRuleID = "021"
const AWSRuleDescription = "CloudFront distribution uses outdated SSL/TLS protocols."
const AWSRelatedService = ["CloudWatch", "CloudFront"]
const AWSRule = "Medium"
const AWSRuleMitigation = "Use the most modern TLS/SSL policies available 
You should not use outdated/insecure TLS versions for encryption. You should be using TLS v1.2+.
"
const AWSRuleItem = `
resource "aws_cloudfront_distribution"  {
  viewer_certificate {
    cloudfront_default_certificate = true
	minimum_protocol_version = "TLSv1.0"
  }
}
`