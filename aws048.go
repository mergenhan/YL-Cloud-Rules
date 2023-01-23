package rules


const AWSRuleID = "048"
const AWSRuleDescription = "EFS Encryption has not been enabled"
const AWSRelatedService = ["EFS"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Enable encryption for EFS
If your organization is subject to corporate or regulatory policies that require encryption of data and metadata at rest, we recommend creating a file system that is encrypted at rest, and mounting your file system using encryption of data in transit.
"
const AWSRuleItem = `
resource "aws_efs_file_system"  {
  name       = "bar"
  encrypted  = false
  kms_key_id = ""
}`