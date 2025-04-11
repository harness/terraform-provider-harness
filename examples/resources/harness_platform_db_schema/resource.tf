resource "harness_platform_db_schema" "test" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
  name       = "name"
  service    = "service1"
  tags       = ["foo:bar", "bar:foo"]
  schema_source {
    connector    = "gitConnector"
    repo         = "TestRepo"
    location     = "db/example-changelog.yaml"
    archive_path = "path/to/archive.zip"
  }
}

resource "harness_platform_db_schema" "test" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
  name       = "name"
  service    = "service1"
  type       = "Repository"
  tags       = ["foo:bar", "bar:foo"]
  schema_source {
    connector    = "gitConnector"
    repo         = "TestRepo"
    location     = "db/example-changelog.yaml"
    archive_path = "path/to/archive.zip"
  }
}


resource "harness_platform_db_schema" "test" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
  name       = "name"
  service    = "service1"
  type       = "Script"
  tags       = ["foo:bar", "bar:foo"]
  changelog_script {
    image    = "plugins/image"
    command  = "echo \\\"hello dbops\\\""
    shell    = "sh/bash"
    location = "db/example-changelog.yaml"
  }
}