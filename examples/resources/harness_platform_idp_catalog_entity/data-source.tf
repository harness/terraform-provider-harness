# Account level catalog entity
data "harness_platform_idp_catalog_entity" "test" {
  identifier = "identifier"
  kind       = "component"
}

# Org level catalog entity
data "harness_platform_idp_catalog_entity" "test" {
  identifier = "identifier"
  kind       = "component"
  org_id     = "org_id"
}

# Project level catalog entity
data "harness_platform_idp_catalog_entity" "test" {
  identifier = "identifier"
  kind       = "component"
  org_id     = "org_id"
  project_id = "project_id"
}
