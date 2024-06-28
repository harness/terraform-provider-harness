resource "harness_platform_db_schema" "test" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
  name       = "name"
  service    = "service1"
  tags       = ["foo:bar", "bar:foo"]
  change_log {
    connector = "gitConnector"
    repo      = "TestRepo"
    location  = "db/example-changelog.yaml"
  }
}