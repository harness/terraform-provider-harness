//Create Project level webhook
resource "harness_platform_gitx_webhook" "test" {
  identifier    = "webhook Identifier"
  name          = "webhook name"
  project_id    = "projectIdentifier"
  org_id        = "orgIdentifier"
  repo_name     = "repo name"
  connector_ref = "connectorRef"
}

//Create Org level webhook
resource "harness_platform_gitx_webhook" "test" {
  identifier    = "webhook Identifier"
  name          = "webhook name"
  org_id        = "orgIdentifier"
  repo_name     = "repo name"
  connector_ref = "connectorRef"
}

//Create Account level webhook
resource "harness_platform_gitx_webhook" "test" {
  identifier    = "webhook Identifier"
  name          = "webhook name"
  repo_name     = "repo name"
  connector_ref = "connectorRef"
}