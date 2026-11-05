data "harness_platform_idp_integration_kind_mapping" "example" {
  integration_id = "servicenow_prod"
  org_id         = "default"
  kind           = "cmdb_ci_service"
}
