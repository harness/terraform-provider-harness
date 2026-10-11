# The streams an Airbyte integration's connector can read, with their sync modes.
data "harness_platform_idp_integration_catalog" "example" {
  integration_id = "servicenow_prod"
  org_id         = "default"
}
