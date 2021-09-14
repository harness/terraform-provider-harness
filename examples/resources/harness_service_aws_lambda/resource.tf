resource "harness_application" "example" {
  name = "example"
}

resource "harness_service_aws_lambda" "example" {
  app_id      = harness_application.example.id
  name        = "my-lambda-service"
  description = "Service for deploying AWS Lambda functions."
}
