// Create file
resource "harness_platform_file_store_node" "file" {
  org_id                    = "org_id"
  project_id                = "project_id"
  identifier                = "identifier"
  name                      = "name"
  description               = "description"
  tags                      = ["foo:bar", "baz:qux"]
  parent_identifier          = "parentIdentifier"
  content                   = "content"
  mime_type                 = "text"
  type                      = "File"
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
  parent_identifier         = "parentIdentifier"
  type                      = "Folder"
}

