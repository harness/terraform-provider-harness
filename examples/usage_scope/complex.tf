resource "harness_application" "example" {
  name = "myapp"
}

resource "harness_environment" "qa" {
  name   = "qa"
  app_id = harness_application.example.id
  type   = "NON_PROD"
}

resource "harness_cloudprovider_kubernetes" "k8s" {
  name = "k8s"

  // Example of scoping to all non-prod environments of a specific application
  usage_scope {
    application_id          = harness_application.example.id
    environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
  }

  // Example of scoping to a specific environment
  usage_scope {
    application_id = harness_application.example.id
    environment_id = harness_environment.qa.id
  }
}
