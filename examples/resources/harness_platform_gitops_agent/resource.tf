resource "harness_platform_gitops_agent" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  name       = "name"
  type       = "MANAGED_ARGO_PROVIDER"
  metadata {
    namespace         = "namespace"
    high_availability = true
  }
}
