# Look up a delegate token at account level by name
data "harness_platform_delegatetoken" "account_level" {
  name       = "account-delegate-token"
  account_id = "account_id"
}

# Look up a delegate token at organization level
data "harness_platform_delegatetoken" "org_level" {
  name       = "org-delegate-token"
  account_id = "account_id"
  org_id     = "org_id"
}

# Look up a delegate token at project level
data "harness_platform_delegatetoken" "project_level" {
  name       = "project-delegate-token"
  account_id = "account_id"
  org_id     = "org_id"
  project_id = "project_id"
}
