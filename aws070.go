package rules

const AWSRuleID = "070"
const AWSRuleDescription = "AWS ES Domain should have logging enabled"
const AWSRelatedService = ["EKS","ES"]
const AWSRule = "Medium"
const AWSRuleMitigation = "Enable logging for ElasticSearch domains
AWS ES domain should have logging enabled by default.
"
const AWSRuleItem = `
resource "aws_elasticsearch_domain" {

  log_publishing_options {
    cloudwatch_log_group_arn = aws_cloudwatch_log_group.example.arn
    log_type                 = "INDEX_SLOW_LOGS"
  }
}
`