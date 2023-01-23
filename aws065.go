package rules

const AWSRuleID = "065"
const AWSRuleDescription = "Cloudtrail should be encrypted at rest to secure access to sensitive trail data"
const AWSRelatedService = ["Cloudtrail"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enable encryption at rest
Cloudtrail logs should be encrypted at rest to secure the sensitive data. Cloudtrail logs record all activity that occurs in the the account through API calls and would be one of the first places to look when reacting to a breach.
"
const AWSRuleItem = `
resource "aws_cloudtrail"  {
  is_multi_region_trail = true

  event_selector {
    read_write_type           = "All"
    include_management_events = true

    data_resource {
      type = "AWS::S3::Object"
      values = ["${data.aws_s3_bucket.important-bucket.arn}/"]
    }
  }
}
`