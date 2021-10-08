resource "harness_cloudprovider_kubernetes" "test" {
  name = "test"

  usage_scope {
    environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
  }

  usage_scope {
    environment_filter_type = "PRODUCTION_ENVIRONMENTS"
  }

}
