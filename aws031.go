package rules

const AWSRuleID = "031"
const AWSRuleDescription = "Elasticsearch domain isn't encrypted at rest."
const AWSRelatedService = ["ES", "CloudWatch"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enable ElasticSearch domain encryption You should ensure your Elasticsearch data is encrypted at rest to help prevent sensitive information from being read by unauthorised users. 
"
const AWSRuleItem = `
resource "aws_elasticsearch_domain"  {
  domain_name = "domain-foo"

  encrypt_at_rest {
    enabled = false
  }
}
`