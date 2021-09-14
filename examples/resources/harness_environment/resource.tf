resource "harness_application" "test" {
  name = "%[1]s"
}

resource "harness_service_kubernetes" "test" {
  app_id       = harness_application.test.id
  name         = "%[1]s"
  helm_version = "V2"
  description  = "description"

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

resource "harness_environment" "test" {
  app_id = harness_application.test.id
  name   = "%[1]s"
  type   = "%[2]s"

  variable_override {
    service_name = harness_service_kubernetes.test.name
    name         = "test"
    value        = "override"
    type         = "TEXT"
  }

  variable_override {
    service_name = harness_service_kubernetes.test.name
    name         = "test2"
    value        = "override2"
    type         = "TEXT"
  }
}
