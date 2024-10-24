resource "harness_platform_infra_module" "example" {
  description          = "example"
  name                 = "name"
  system               = "provider"
  repository           = "https://github.com/org/repo"
  repository_branch    = "main"
  repository_path      = "tf/aws/basic"
  repository_connector = harness_platform_connector_github.test.id
}
