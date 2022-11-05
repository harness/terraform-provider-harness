resource "harness_platform_connector_kubernetes_cloud_cost" "example" {
  identifier  = "identifier"
  name        = "name"
  description = "example"
  tags        = ["foo:bar"]

  features_enabled = ["VISIBILITY", "OPTIMIZATION"]
  connector_ref    = "connector_ref"
}
