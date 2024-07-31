//Create Project level webhook
resource "harness_platform_gitx_webhook" "test" {
  identifier = "webhook Identifier"
  name       = "webhook name"
  project_id = "projectIdentifier"
  org_id     = "orgIdentifier"
}

//Create Org level webhook
resource "harness_platform_gitx_webhook" "test" {
  identifier = "webhook Identifier"
  name       = "webhook name"
  org_id     = "orgIdentifier"
}

//Create Account level webhook
resource "harness_platform_gitx_webhook" "test" {
  identifier = "webhook Identifier"
  name       = "webhook name"
}