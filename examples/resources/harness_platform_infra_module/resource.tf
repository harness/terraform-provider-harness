resource "harness_platform_infra_module" "example" {
  description          = "example"
  name                 = "name"
  system               = "provider"
  repository           = "repo"
  repository_branch    = "main"
  repository_path      = "tf/aws/basic"
  repository_connector = harness_platform_connector_github.test.id
  onboarding_pipeline         = "my_onboarding_pipeline"
  onboarding_pipeline_org     = "default"
  onboarding_pipeline_project = "IaCM_Project"
  onboarding_pipeline_sync    = true
  storage_type                = "artifact"
  connector_org               = "default"
  connector_project           = "my_project"
}
