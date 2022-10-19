data "harness_platform_gitops_agent" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  name       = "name"
  type       = "CONNECTED_ARGO_PROVIDER"
  metadata {
    namespace         = "test"
    high_availability = true
  }
}
