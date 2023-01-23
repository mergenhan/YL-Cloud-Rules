package rules
const AWSRuleID = "044"
const AWSRuleDescription = "AWS provider has access credentials specified."
const AWSRelatedService = ["IAM"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Don't include access credentials in plain text
The AWS provider block should not contain hardcoded credentials. These can be passed in securely as runtime using environment variables.
"
const AWSRuleItem = `
provider "aws" {
  access_key = string
  secret_key = string
}