package rules

const AWSRuleID = "011"
const AWSRuleDescription = "A database resource is marked as publicly accessible."
const AWSRelatedService = ["EC2", "RDS", "CloudFront"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Set the database to not be publically accessible 
Database resources should not publicly available. You should limit all access to the minimum that is required for your application to function. 
"
const AWSRuleItem = `
resource "aws_db_instance"   {
	publicly_accessible = true
}