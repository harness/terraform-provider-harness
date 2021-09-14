resource "harness_application" "example" {
  name = "example"
}

resource "harness_service_tanzu" "example" {
  app_id      = harness_application.example.id
  name        = "tanzu-svc"
  description = "A service for deploying Tanzu applications."
}
