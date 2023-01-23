package rules


const AWSRuleID = "013"
const AWSRuleDescription = "Task definition defines sensitive environment variable(s)."
const AWSRelatedService = ["EC2", "RDS"]
const AWSRule = "Critical"
const AWSRuleMitigation = "Use secrets for the task definition You should not make secrets available to a user in plaintext in any scenario. Secrets can instead be pulled from a secure secret storage system by the service requiring them.  
"
const AWSRuleItem = `
resource "aws_ecs_task_definition"  {
  container_definitions = <<EOF
[
  {
    "name": "my_service",
    "essential": true,
    "memory": 256,
    "environment": [
      { "name": "ENVIRONMENT", "value": "development" },
      { "name": "DATABASE_PASSWORD", "value": "No"}
    ]
  }
]
EOF

}