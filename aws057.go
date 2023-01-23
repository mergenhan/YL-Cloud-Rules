package rules

const AWSRuleID = "057"
const AWSRuleDescription = "Domain logging should be enabled for Elastic Search domains"
const AWSRelatedService = ["ES"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enable logging for ElasticSearch domains
Amazon ES exposes four Elasticsearch logs through Amazon CloudWatch Logs: error logs, search slow logs, index slow logs, and audit logs. 

Search slow logs, index slow logs, and error logs are useful for troubleshooting performance and stability issues. 

Audit logs track user activity for compliance purposes. 

All the logs are disabled by default. 
"
const AWSRuleItem = `
resource "aws_elasticsearch_domain"  {
  domain_name           = "example"
  elasticsearch_version = "1.5"
}

resource "aws_elasticsearch_domain"  {
  domain_name           = "example"
  elasticsearch_version = "1.5"

  log_publishing_options {
    cloudwatch_log_group_arn = aws_cloudwatch_log_group.example.arn
    log_type                 = "INDEX_SLOW_LOGS"
    enabled                  = false  
  }
}