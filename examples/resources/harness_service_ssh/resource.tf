resource "harness_application" "example" {
  name = "example"
}

resource "harness_service_ssh" "example" {
  app_id        = harness_application.example.id
  artifact_type = "TAR"
  name          = "ssh-example"
  description   = "Service for deploying applications with SSH."
}
