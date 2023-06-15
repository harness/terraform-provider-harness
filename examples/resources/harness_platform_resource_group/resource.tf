resource "harness_platform_resource_group" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  account_id           = "account_id"
  allowed_scope_levels = ["account"]
  included_scopes {
    filter     = "EXCLUDING_CHILD_SCOPES"
    account_id = "account_id"
  }
  resource_filter {
    include_all_resources = false
    resources {
      resource_type = "CONNECTOR"
      attribute_filter {
        attribute_name   = "category"
        attribute_values = ["CLOUD_COST"]
      }
    }
  }
}
