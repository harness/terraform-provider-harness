resource "harness_platform_repo" "test" {
  identifier         = "test_repo_123"
  name               = "test_repo_123"
  org_identifier     = "test_org_123"
  project_identifier = "test_project_123"
  default_branch     = "main"
  description        = "test_description_123"
  account_id         = "account_id"
  provider_repo      = "octocat/hello-worId"
  is_public          = true
  type               = "github"
}
