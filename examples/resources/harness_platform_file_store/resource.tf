// Create file
resource "harness_platform_file_store_node" "file" {
  org_id                    = "org_id"
  project_id                = "project_id"
  identifier                = "identifier"
  name                      = "name"
  description               = "description"
  tags                      = ["foo:bar", "baz:qux"]
  parent_identifier         = "parent_identifier"
  file_content_path         = "file_content_path"
  mime_type                 = "mime_type"
  type                      = "FILE"
  file_usage                = "ManifestFile|Config|Script"
}

// Create folder
resource "harness_platform_file_store_node" "folder" {
  org_id                    = "org_id"
  project_id                = "project_id"
  identifier                = "identifier"
  name                      = "name"
  description               = "description"
  tags                      = ["foo:bar", "baz:qux"]
  parent_identifier         = "parent_identifier"
  type                      = "FOLDER"
}

