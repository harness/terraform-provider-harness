data "harness_platform_usergroup" "example_by_id" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
}
data "harness_platform_environment" "example_by_name" {
  name       = "name"
  org_id     = "org_id"
  project_id = "project_id"
}
