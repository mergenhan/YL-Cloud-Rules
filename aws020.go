package rules

const AWSRuleID = "020"
const AWSRuleDescription = "CloudFront distribution allows unencrypted (HTTP) communications."
const AWSRelatedService = ["CloudFront"]
const AWSRule = "Critical"
const AWSRuleMitigation =  "Only allow HTTPS for CloudFront distribution communication 
Plain HTTP is unencrypted and human-readable. This means that if a malicious actor was to eavesdrop on your connection, they would be able to see all of your data flowing back and forth.

You should use HTTPS, which is HTTP over an encrypted (TLS) connection, meaning eavesdroppers cannot read your traffic.
"
const AWSRuleItem = `
resource "aws_cloudfront_distribution"   {
	default_cache_behavior {
	    viewer_protocol_policy = "allow-all"
	  }
}
`