package rules

const AWSRuleID = "062"
const AWSRuleDescription = "User data for EC2 instances must not contain sensitive AWS keys"
const AWSRelatedService = ["EC2"]
const AWSRule = "High"
const AWSRuleItem = `
resource "aws_instance"  {

  ami           = "ami-12345667"
  instance_type = "t2.small"

  user_data = <<EOF
export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
export AWS_DEFAULT_REGION=us-west-2 
EOF
}
`