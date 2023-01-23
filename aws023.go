package rules

const AWSRuleID = "023"
const AWSRuleDescription = "ECR repository has image scans disabled."
const AWSRelatedService = ["ECR"]
const AWSRule = "High"
const AWSRuleMitigation = "Enable ECR image scanning 
Repository image scans should be enabled to ensure vulnerable software can be discovered and remediated as soon as possible.
"
const AWSRuleItem = `
resource "aws_ecr_repository"  {
  name                 = "bar"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = false
  }
}
`