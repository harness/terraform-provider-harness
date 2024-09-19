resource "harness_governance_rule" "example" {
  identifier         = "identifier"
  name               = "name"
  cloud_provider     = "AWS/AZURE/GCP"
  description        = "description"
  rules_yaml 		     = "policies:\n  - name: aws-list-ec2\n    resource: aws.ec2"
}