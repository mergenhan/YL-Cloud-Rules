package rules

const AWSRuleID = "059"
const AWSRuleDescription = "Athena databases and workgroup configurations are created unencrypted at rest by default, they should be encrypted"
const AWSRelatedService = ["Athena", "Neptün", "Lambda"]
const AWSRule = "Critical"

const AWSRuleItem = `
resource "aws_athena_database"  {
  name   = "database_name"
  bucket = aws_s3_bucket.hoge.bucket
}

resource "aws_athena_workgroup"  {
  name = "example"

  configuration {
    enforce_workgroup_configuration    = true
    publish_cloudwatch_metrics_enabled = true

    result_configuration {
      output_location = "s3://${aws_s3_bucket.example.bucket}/output/"
    }
  }
}
`