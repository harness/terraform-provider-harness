data "harness_platform_project" "example_by_id" {
  identifier = "identifier"
  org_id     = "org_id"
}

data "harness_platform_project" "example_by_name" {
  name   = "name"
  org_id = "org_id"
}
