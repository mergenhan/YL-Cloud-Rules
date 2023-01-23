package rules

const AWSRuleID = "032"
const AWSRuleDescription = "Elasticsearch domain uses plaintext traffic for node to node communication."
const AWSRelatedService = ["Kinesis", "ES"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enable encrypted node to node communication Traffic flowing between Elasticsearch nodes should be encrypted to ensure sensitive data is kept private.
"
const AWSRuleItem = `
resource "aws_elasticsearch_domain" {
  domain_name = "domain-foo"

  node_to_node_encryption {
    enabled = false
  }
}