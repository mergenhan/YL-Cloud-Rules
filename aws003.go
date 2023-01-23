package rules

const AWSRuleID = "003"
const AWSRuleDescription = "AWS Classic resource usage."
const AWSRelatedService = ["VPC", "EC2"]
const AWSRuleImpact = "Low"
const AWSRuleMitigation = "Switch to VPC resources"

const AWSRuleItem = `
esource "aws_msk_cluster"  {
	encryption_info {
		encryption_in_transit {
			client_broker = "TLS_PLAINTEXT"
			in_cluster = true
		}
	}
}`