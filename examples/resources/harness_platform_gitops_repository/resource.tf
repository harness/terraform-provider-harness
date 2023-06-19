// Create a git repository at project level
resource "harness_platform_gitops_repository" "example" {
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

// Create a HELM repository at project level
resource "harness_platform_gitops_repository" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"
  repo {
    repo            = "https://charts.helm.sh/stable"
    name            = "repo_name"
    insecure        = true
    connection_type = "HTTPS_ANONYMOUS"
    type_           = "helm"
  }
  upsert = true
}

// Create a OCI HELM repository at project level
resource "harness_platform_gitops_repository" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"
  repo {
    repo            = "ghcr.io/wings-software"
    name            = "repo_name"
    insecure        = false
    username        = "username"
    password        = "ghp_xxxxxxxx"
    connection_type = "HTTPS"
    type_           = "helm"
    enable_oci      = true
  }
  upsert = true
}
