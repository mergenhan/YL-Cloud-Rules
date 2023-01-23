package rules

const AWSRuleID = "052"
const AWSRuleDescription = "RDS encryption has not been enabled at a DB Instance level."
const AWSRelatedService = ["RDS"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enable encryption for RDS clusters and instances
Encryption should be enabled for an RDS Database instances. 

When enabling encryption by setting the kms_key_id. 
"
const AWSRuleItem = `
resource "aws_db_instance"  {
	
}'