resource "harness_platform_variables" "test" {
  identifier = "identifier"
  name       = "name"
  org_id     = "org_id"
  project_id = "project_id"
  type       = "String"
  spec {
    value_type  = "FIXED"
    fixed_value = "fixedValue"
  }
}
