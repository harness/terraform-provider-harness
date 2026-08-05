resource "harness_platform_db_instance" "test" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
  name       = "name"
  tags       = ["foo:bar", "bar:foo"]
  schema     = "schema1"
  branch     = "main"
  connector  = "jdbcConnector"
  context    = "ctx"

  substitute_properties = {
    "key1" = "value1"
    "key2" = "value2"
  }
}

# Using commit_sha to pin changelog to a specific revision
resource "harness_platform_db_instance" "with_commit_sha" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
  name       = "name"
  tags       = ["foo:bar", "bar:foo"]
  schema     = "schema1"
  commit_sha = "abc123def456"
  connector  = "jdbcConnector"
  context    = "ctx"
}

# Using git_tag to pin changelog to a specific tagged revision
resource "harness_platform_db_instance" "with_git_tag" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
  name       = "name"
  tags       = ["foo:bar", "bar:foo"]
  schema     = "schema1"
  git_tag    = "v1.0.0"
  connector  = "jdbcConnector"
  context    = "ctx"
}