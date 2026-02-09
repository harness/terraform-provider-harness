# Account level catalog entity
resource "harness_platform_idp_catalog_entity" "test" {
  identifier = "identifier"
  kind       = "component"
  yaml       = <<-EOT
    apiVersion: harness.io/v1
    kind: Component
    name: Example Catalog
    identifier: "%[1]s"
    type: service
    owner: user:account/admin@harness.io
    spec:
        lifecycle: prod
    metadata:
        tags:
            - test
    EOT
}

# Org level catalog entity
resource "harness_platform_idp_catalog_entity" "test" {
  identifier = "identifier"
  kind       = "component"
  org_id     = "org_id"
  yaml       = <<-EOT
    apiVersion: harness.io/v1
    kind: Component
    name: Example Catalog
    identifier: "identifier"
    orgIdentifier: org_id
    type: service
    owner: user:account/admin@harness.io
    spec:
        lifecycle: prod
    metadata:
        tags:
            - test
    EOT
}

# Project level catalog entity
resource "harness_platform_idp_catalog_entity" "test" {
  identifier = "identifier"
  kind       = "component"
  org_id     = "org_id"
  project_id = "project_id"
  yaml       = <<-EOT
    apiVersion: harness.io/v1
    kind: Component
    name: Example Catalog
    identifier: "identifier"
    orgIdentifier: org_id
    projectIdentifier: project_id
    type: service
    owner: user:account/admin@harness.io
    spec:
        lifecycle: prod
    metadata:
        tags:
            - test
    EOT
}

# Project level catalog entity with git experience
resource "harness_platform_idp_catalog_entity" "test_with_git_experience" {
  identifier = "identifier"
  kind       = "component"
  org_id     = "org_id"
  project_id = "project_id"
  git_details {
    branch_name = "main"
    file_path   = "path/to/file.yaml"
    repo_name   = "repo_name"
    store_type  = "REMOTE"
  }
  yaml = <<-EOT
    apiVersion: harness.io/v1
    kind: Component
    name: Example Catalog
    identifier: "identifier"
    orgIdentifier: org_id
    projectIdentifier: project_id
    type: service
    owner: user:account/admin@harness.io
    spec:
        lifecycle: prod
    metadata:
        tags:
            - test
    EOT
}

# Project level catalog entity with git experience - With import
resource "harness_platform_idp_catalog_entity" "test_with_git_import" {
  identifier      = "identifier"
  org_id          = "org_id"
  project_id      = "project_id"
  import_from_git = true
  git_details {
    branch_name = "main"
    file_path   = "path/to/file.yaml"
    repo_name   = "repo_name"
    store_type  = "REMOTE"
  }
}
