package rules

const AWSRuleID = "014"
const AWSRuleDescription = "Launch configuration with unencrypted block device."
const AWSRelatedService = ["EC2", "EBS"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Turn on encryption for all block devices 
Blocks devices should be encrypted to ensure sensitive data is hel securely at rest.
"
const AWSRuleItem = `
resource "aws_launch_configuration"  {
	root_block_device {
		encrypted = false
	}
}