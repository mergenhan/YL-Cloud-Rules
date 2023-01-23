package rules

const AWSRuleID = "071"
const AWSRuleDescription = "Cloudfront distribution should have Access Logging configured"
const AWSRelatedService = ["Cloudfront"]
const const AWSRule = "Medium"
const AWSRuleMitigation = "Enable logging for CloudFront distributions
You should configure CloudFront Access Logging to create log files that contain detailed information about every user request that CloudFront receives
"
const AWSRuleItem = `
resource "aws_cloudfront_distribution"  {
	// other config
	// no logging_config
}
`