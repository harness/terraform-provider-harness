# Look up a Split environment by name within the FME workspace for a Harness org and project.
data "harness_fme_environment" "production" {
  org_id     = "organization_id"
  project_id = "project_id"
  name       = "Production"
}
