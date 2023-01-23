package rules

const AWSRuleID = "012"
const AWSRuleDescription = "A resource has a public IP address."
const AWSRelatedService = ["EC2", "API Gateway", "S3"]
const AWSRule = "Critical"
const AWSRuleMitigation ="Set the instance to not be publically accessible 
You should limit the provision of public IP addresses for resources. Resources should not be exposed on the public internet, but should have access limited to consumers required for the function of your application. 
"
const AWSRuleItem = `
resource "aws_launch_configuration"   {
	associate_public_ip_address = true
}