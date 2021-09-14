resource "harness_application" "example" {
  name = "example"
}

resource "harness_service_kubernetes" "example" {
  app_id       = harness_application.example.id
  name         = "k8s-svc"
  helm_version = "V3"
  description  = "Service for deploying Kubernetes manifests"

  variable {
    name  = "test"
    value = "test_value"
    type  = "TEXT"
  }

  variable {
    name  = "test2"
    value = "test_value2"
    type  = "TEXT"
  }
}
