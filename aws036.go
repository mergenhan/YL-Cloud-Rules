package rules

const AWSRuleID = "036"
const AWSRuleDescription = "Elasticache Replication Group uses unencrypted traffic."
const AWSRelatedService = ["ES"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enable in transit encryptuon for replication group Traffic flowing between Elasticache replication nodes should be encrypted to ensure sensitive data is kept private.
"
const AWSRuleItem = `
resource "aws_elasticache_replication_group"  {
        replication_group_id = "foo"
        replication_group_description = "my foo cluster"

        transit_encryption_enabled = false
}
`