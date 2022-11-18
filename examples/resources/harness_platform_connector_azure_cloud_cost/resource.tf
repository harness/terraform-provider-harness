resource "harness_platform_connector_azure_cloud_cost" "example" {
  identifier  = "identifier"
  name        = "name"
  description = "example"
  tags        = ["foo:bar"]

  features_enabled = ["BILLING", "VISIBILITY", "OPTIMIZATION"]
  tenant_id        = "tenant_id"
  subscription_id  = "subscription_id"
  billing_export_spec {
    storage_account_name = "storage_account_name"
    container_name       = "container_name"
    directory_name       = "directory_name"
    report_name          = "report_name"
    subscription_id      = "subscription_id"
  }
}
