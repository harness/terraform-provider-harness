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
