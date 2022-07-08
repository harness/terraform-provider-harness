resource "harness_platform_roles" "example" {
  identifier           = "identifier"
  name                 = "name"
  description          = "test"
  tags                 = ["foo:bar"]
  permissions          = ["core_pipeline_edit"]
  allowed_scope_levels = ["project"]
}
