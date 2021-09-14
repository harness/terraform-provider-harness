resource "harness_application" "example" {
  name = "example"
}

resource "harness_service_aws_codedeploy" "example" {
  app_id      = harness_application.example.id
  name        = "aws-codedeploy"
  description = "Service for AWS codedeploy applications."
}
