package rules


const AWSRuleID = "061"
const AWSRuleDescription = "API Gateway stages for V1 and V2 should have access logging enabled"
const AWSRelatedService = ["AWS Gateway"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enable logging for API Gateway stages
API Gateway stages should have access log settings block configured to track all access to a particular stage. This should be applied to both v1 and v2 gateway stages.
"
const AWSRuleItem = `
resource "aws_apigatewayv2_stage"  {
  api_id = aws_apigatewayv2_api.example.id
  name   = "example-stage"
}

resource "aws_api_gateway_stage" {
  deployment_id = aws_api_gateway_deployment.example.id
  rest_api_id   = aws_api_gateway_rest_api.example.id
  stage_name    = "example"
}
`