package rules

const AWSRuleID = "010"
const AWSRuleDescription = "An outdated SSL policy is in use by a load balancer."
const AWSRelatedService = ["EC2", "ELB", "CloudFront"]
const AWSRule = "Medium"
const AWSRuleMitigation = "Use a more recent TLS/SSL policy for the load balancer 
You should not use outdated/insecure TLS versions for encryption. You should be using TLS v1.2+. 
"
const AWSRuleItem = `
resource "aws_alb_listener"  {
	ssl_policy = "ELBSecurityPolicy-TLS-1-1-2017-01"
	protocol = "HTTPS"
}