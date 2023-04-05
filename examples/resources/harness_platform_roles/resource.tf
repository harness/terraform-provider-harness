resource "harness_platform_roles" "example" {
  identifier           = "identifier"
  name                 = "name"
  description          = "test"
  tags                 = ["foo:bar"]
  permissions          = ["core_resourcegroup_view"]
  allowed_scope_levels = ["account"]
}
