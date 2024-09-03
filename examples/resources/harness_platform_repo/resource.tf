resource "harness_platform_repo" "test" {
  identifier     = "test_repo_123"
  name           = "test_repo_123"
  org_id         = "test_org_123"
  project_id     = "test_project_123"
  default_branch = "main"
  description    = "test_description_123"
  source {
    repo = "octocat/hello-worId"
    type = "github"
  }
}
