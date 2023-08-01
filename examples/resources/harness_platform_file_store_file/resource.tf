// Create file
resource "harness_platform_file_store_file" "example" {
  org_id            = "org_id"
  project_id        = "project_id"
  identifier        = "identifier"
  name              = "name"
  description       = "description"
  tags              = ["foo:bar", "baz:qux"]
  parent_identifier = "parent_identifier"
  file_content_path = "file_content_path"
  mime_type         = "mime_type"
  file_usage        = "MANIFEST_FILE|CONFIG|SCRIPT"
}

