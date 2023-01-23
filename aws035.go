package rules

const AWSRuleID = "035"
const AWSRuleDescription = "Unencrypted Elasticache Replication Group."
const AWSRelatedService = [ "ES"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enable encryption for replication group You should ensure your Elasticache data is encrypted at rest to help prevent sensitive information from being read by unauthorised users.
"
const AWSRuleItem = `
resource "aws_elasticache_replication_group"  {
        replication_group_id = "foo"
        replication_group_description = "my foo cluster"

        at_rest_encryption_enabled = false
}