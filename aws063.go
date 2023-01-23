package rules

const AWSRuleID = "063"
const AWSRuleDescription = "Cloudtrail should be enabled in all regions regardless of where your AWS resources are generally homed"
const AWSRelatedService = ["Cloudtrail"]
const AWSRule = "High"
const AWSRuleMitigation = "Enable Cloudtrail in all regions
When creating Cloudtrail in the AWS Management Console the trail is configured by default to be multi-region, this isn't the case with the Terraform resource. Cloudtrail should cover the full AWS account to ensure you can track changes in regions you are not actively operting in.
"
const AWSRuleItem = `
resource "aws_cloudtrail"  {
  event_selector {
    read_write_type           = "All"
    include_management_events = true

    data_resource {
      type = "AWS::S3::Object"
      values = ["${data.aws_s3_bucket.important-bucket.arn}/"]
    }
  }
}'