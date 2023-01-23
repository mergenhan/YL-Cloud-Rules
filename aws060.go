package rules

const AWSRuleID = "060"
const AWSRuleDescription = "Athena workgroups should enforce configuration to prevent client disabling encryption"
const AWSRelatedService = ["Athena", "Neptün", "Lambda"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enforce the configuration to prevent client overrides
Athena workgroup configuration should be enforced to prevent client side changes to disable encryption settings.
"
const AWSRuleItem = `
resource "aws_athena_workgroup"  {
  name = "example"

  configuration {
    enforce_workgroup_configuration    = false
    publish_cloudwatch_metrics_enabled = true

    result_configuration {
      output_location = "s3://${aws_s3_bucket.example.bucket}/output/"

      encryption_configuration {
        encryption_option = "SSE_KMS"
        kms_key_arn       = aws_kms_key.example.arn
      }
    }
  }
}

resource "aws_athena_workgroup" "bad_example" {
  name = "example"

}
`