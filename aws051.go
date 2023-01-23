package rules


const AWSRuleID = "051"
const AWSRuleDescription = "There is no encryption specified or encryption is disabled on the RDS Cluster."
const AWSRelatedService = ["RDS"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enable encryption for RDS clusters and instances
Encryption should be enabled for an RDS Aurora cluster. 

When enabling encryption by setting the kms_key_id, the storage_encrypted must also be set to true. 
"
const AWSRuleItem = `
resource "aws_rds_cluster"  {
  name       = "bar"
  kms_key_id = ""
}`