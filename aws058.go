package rules

const AWSRuleID = "058"
const AWSRuleDescription = "Ensure that lambda function permission has a source arn specified"
const AWSRelatedService = ["ES","Lambda"]
const AWSRule = "High"
const AWSRuleMitigation = "Always provide a source arn for Lambda permissions When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to. 

Without this, any resource from principal will be granted permission – even if that resource is from another account. 

For S3, this should be the ARN of the S3 Bucket. For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule. For API Gateway, this should be the ARN of the API"

const AWSRuleItem = `
resource "aws_lambda_permission"  {
  statement_id  = "AllowExecutionFromSNS"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.func.function_name
  principal     = "sns.amazonaws.com"
}
`