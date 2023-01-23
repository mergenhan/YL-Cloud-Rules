package rules

const AWSRuleID = "022"
const AWSRuleDescription = "A MSK cluster allows unencrypted data in transit."
const AWSRelatedService = ["MSK"]
const AWSRule = "Critical"
const AWSRuleMitigation =  "Enable in transit encryption 
Encryption should be forced for Kafka clusters, including for communication between nodes. This ensure sensitive data is kept private.
"
const AWSRuleItem = `
resource "aws_msk_cluster"  {
	encryption_info {
		encryption_in_transit {
			client_broker = "TLS_PLAINTEXT"
			in_cluster = true
		}
	}
}
`