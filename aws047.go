package rules

const AWSRuleID = "047"
const AWSRuleDescription = "AWS SQS policy document has wildcard action statement."
const AWSRelatedService = ["SQS"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Keep policy scope to the minimum that is required to be effective
SQS Policy actions should always be restricted to a specific set.

This ensures that the queue itself cannot be modified or deleted, and prevents possible future additions to queue actions to be implicitly allowed.
"
const AWSRuleItem = `
resource "aws_sqs_queue_policy"  {
  queue_url = aws_sqs_queue.q.id

  policy = <<POLICY
{
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": "*",
      "Action": "*"
    }
  ]
}
POLICY
}
`