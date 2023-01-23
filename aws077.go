package rules

const AWSRuleID = "077"
const AWSRuleDescription = "S3 Data should be versioned"
const AWSRelatedService = ["S3"]
const AWSRule = "Low"
const AWSRuleMitigation = "Enable versioning to protect against accidental/malicious removal or modification
Versioning in Amazon S3 is a means of keeping multiple variants of an object in the same bucket. 
You can use the S3 Versioning feature to preserve, retrieve, and restore every version of every object stored in your buckets. 
With versioning you can recover more easily from both unintended user actions and application failures.
"
const AWSRuleItem = `
resource "aws_s3_bucket" {

}
`