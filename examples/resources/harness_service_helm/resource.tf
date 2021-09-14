resource "harness_application" "example" {
  name = "example"
}

resource "harness_service_helm" "example" {
  app_id      = harness_application.example.id
  name        = "helm-example-service"
  description = "Service for deploying native Helm application.s"
}
