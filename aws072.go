package rules

const AWSRuleID = "072"
const AWSRuleDescription = "Viewer Protocol Policy in Cloudfront Distribution Cache should always be set to HTTPS"
const AWSRelatedService = ["Cloudfront"]
const AWSRule = "Medium"
const AWSRuleMitigation = "Only use HTTPS in the Viewer Protocol Policy
CloudFront connections should be encrypted during transmission over networks that can be accessed by malicious individuals. 
A CloudFront distribution should only use HTTPS or Redirect HTTP to HTTPS for communication between viewers and CloudFront.
"
const AWSRuleItem = `
resource "aws_cloudfront_distribution" {
	// other cloudfront distribution config

	default_cache_behavior {
		// other cache config

		viewer_protocol_policy = "allow-all" // including HTTP
	}
}

resource "aws_cloudfront_distribution"  {
	// other cloudfront distribution config

	default_cache_behavior {
		// other cache config

		viewer_protocol_policy = "https-only" // HTTPS by default...
	}

	ordered_cache_behavior {
		// other cache config

		viewer_protocol_policy = "allow-all" // ...but HTTP for some other resources
	}
}
`