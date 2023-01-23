package rules

const AWSRuleID = "018"
const AWSRuleDescription = "Missing description for security group/security group rule."
const AWSRelatedService = ["EC2", "VPC"]
const AWSRule = "Critical"
const AWSRuleMitigation ="Add descriptions for all security groups anf rules 
Security groups and security group rules should include a description for auditing purposes.

Simplifies auditing, debugging, and managing security groups.
"
const AWSRuleItem = `
resource "aws_security_group" {
  name        = "http"

  ingress {
    description = "HTTP from VPC"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = [aws_vpc.main.cidr_block]
  }
}