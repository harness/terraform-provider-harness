resource "harness_platform_db_instance" "test" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
  name       = "name"
  tags       = ["foo:bar", "bar:foo"]
  schema = "schema1"
  branch = "main"
  connector = "jdbcConnector"
  context = "ctx"
}