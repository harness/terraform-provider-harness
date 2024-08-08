data "harness_platform_gitops_app_project_mapping" "example" {
  account_id = "account_id"
  org_id     = "organization_id"
  project_id = "project_id"
  agent_id   = "agent_id"
  argo_proj_name = "argo_proj_name"
}