resource "harness_application" "example" {
  name = "example"
}

resource "harness_service_ami" "example" {
  app_id      = harness_application.example.id
  name        = "ami-example"
  description = "Service for deploying AMI's"
}
