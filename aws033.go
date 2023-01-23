package rules

const AWSRuleID = "033"
const AWSRuleDescription = "Elasticsearch doesn't enforce HTTPS traffic."
const AWSRelatedService = ["Kinesis", "ES"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enforce the use of HTTPS for ElasticSearch Plain HTTP is unencrypted and human-readable. This means that if a malicious actor was to eavesdrop on your connection, they would be able to see all of your data flowing back and forth.

You should use HTTPS, which is HTTP over an encrypted (TLS) connection, meaning eavesdroppers cannot read your traffic.
"
const AWSRuleItem = `
resource "aws_elasticsearch_domain"  {
  domain_name = "domain-foo"

  domain_endpoint_options {
    enforce_https = false
  }
}