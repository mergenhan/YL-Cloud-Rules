package rules

const AWSRuleID = "053"
const AWSRuleDescription = "Encryption for RDS Perfomance Insights should be enabled."
const AWSRelatedService = ["RDS"]
const AWSRule = "Critical"

const AWSRuleItem = `
resource "aws_rds_cluster_instance"  {
  name                 = "bar"
  performance_insights_enabled = true
  performance_insights_kms_key_id = ""
}
`