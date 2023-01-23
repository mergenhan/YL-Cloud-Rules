package rules

const AWSRuleID = "064"
const AWSRuleDescription = "Cloudtrail log validation should be enabled to prevent tampering of log data"
const AWSRelatedService = ["Cloudtrail"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Turn on log validation for Cloudtrail
Log validation should be activated on Cloudtrail logs to prevent the tampering of the underlying data in the S3 bucket. It is feasible that a rogue actor compromising an AWS account might want to modify the log data to remove trace of their actions.
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