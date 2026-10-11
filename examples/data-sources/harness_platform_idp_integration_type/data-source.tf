# The kinds an integration type syncs and the fields each kind can map.
data "harness_platform_idp_integration_type" "example" {
  type = "ServiceNow"
}

output "kinds" {
  value = data.harness_platform_idp_integration_type.example.kinds[*].kind
}
