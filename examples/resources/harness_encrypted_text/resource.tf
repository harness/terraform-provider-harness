data "harness_secret_manager" "default" {
  default = true
}

resource "harness_encrypted_text" "example" {
  name              = "example-secret"
  value             = "someval"
  secret_manager_id = data.harness_secret_manager.default.id

  usage_scope {
    environment_filter_type = "PRODUCTION_ENVIRONMENTS"
  }

  usage_scope {
    environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
  }
}
