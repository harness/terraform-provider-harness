data "harness_platform_gitops_repository" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"
  repo {
    repo            = "https://github.com/willycoll/argocd-example-apps.git"
    name            = "repo_name"
    insecure        = true
    connection_type = "HTTPS_ANONYMOUS"
  }
  upsert = true
}