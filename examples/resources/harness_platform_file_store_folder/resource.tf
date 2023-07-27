// Create folder
resource "harness_platform_file_store_folder" "example" {
  org_id            = "org_id"
  project_id        = "project_id"
  identifier        = "identifier"
  name              = "name"
  description       = "description"
  tags              = ["foo:bar", "baz:qux"]
  parent_identifier = "parent_identifier"
}

