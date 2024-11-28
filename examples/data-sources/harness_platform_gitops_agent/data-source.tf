data "harness_platform_gitops_agent" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
}

# setting with_credentials to true returns the stored token for the agent.
# This is populated in the data source's output only if the agent is not currently authenticated to harness.
# Setting this to true also correctly populates the is_authenticated field in response
data "harness_platform_gitops_agent" "example" {
  identifier       = "identifier"
  account_id       = "account_id"
  project_id       = "project_id"
  org_id           = "org_id"
  with_credentials = true
}

