# Account level setting
resource "harness_platform_setting" "account" {
  identifier      = "enforce_git_experience"
  value           = "true"
  allow_overrides = true
}

# Organization level setting that projects cannot override
resource "harness_platform_setting" "org" {
  identifier      = "git_experience_repo_allowlist"
  org_id          = "org_id"
  value           = "my-org/allowed-repo"
  allow_overrides = false
}

# Project level settings, configured when the project is created
resource "harness_platform_project" "example" {
  identifier = "example"
  name       = "example"
  org_id     = "org_id"
}

locals {
  default_settings = {
    default_connector_for_git_experience = "account.gitlab_connector"
    default_repo_for_git_experience      = "my-org/my-repo"
  }
}

resource "harness_platform_setting" "project" {
  for_each = local.default_settings

  identifier = each.key
  org_id     = harness_platform_project.example.org_id
  project_id = harness_platform_project.example.id
  value      = each.value
}
