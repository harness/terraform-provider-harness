resource "harness_platform_template_filters" "test" {
  identifier = "identifier"
  name       = "name"
  org_id     = "org_id"
  project_id = "project_id"
  type       = "Template"
  filter_properties {
    tags        = ["foo:bar"]
    filter_type = "Template"
  }
  filter_visibility = "EveryOne"
}
