package rules

const AWSRuleID = "034"
const AWSRuleDescription = "Elasticsearch domain endpoint is using outdated TLS policy."
const AWSRelatedService = ["Kinesis", "ES",]
const AWSRule = "Critical"
const AWSRuleMitigation = "Use the most modern TLS/SSL policies available You should not use outdated/insecure TLS versions for encryption. You should be using TLS v1.2+.
"
const AWSRuleItem = `
resource "aws_elasticsearch_domain"  {
  domain_name = "domain-foo"

  domain_endpoint_options {
    enforce_https = true
    tls_security_policy = "Policy-Min-TLS-1-0-2019-07"
  }
}