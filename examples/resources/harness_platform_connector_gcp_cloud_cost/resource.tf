resource "harness_platform_connector_gcp_cloud_cost" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  features_enabled      = ["BILLING", "VISIBILITY", "OPTIMIZATION"]
  gcp_project_id        = "gcp_project_id"
  service_account_email = "service_account_email"
  billing_export_spec {
    data_set_id = "data_set_id"
    table_id    = "table_id"
  }
}
