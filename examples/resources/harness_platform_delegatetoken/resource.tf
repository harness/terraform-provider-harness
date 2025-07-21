# Create delegate token at account level
resource "harness_platform_delegatetoken" "account_level" {
  name = "account-delegate-token"
  account_id = "account_id"
}

# Create delegate token at organization level
resource "harness_platform_delegatetoken" "org_level" {
  name = "org-delegate-token"
  account_id = "account_id"
  org_id = "org_id"
}

# Create delegate token at project level
resource "harness_platform_delegatetoken" "project_level" {
  name = "project-delegate-token"
  account_id = "account_id"
  org_id = "org_id"
  project_id = "project_id"
}

# Create delegate token with auto-expiry
resource "harness_platform_delegatetoken" "expiry_token" {
  name = "expiry-delegate-token"
  account_id = "account_id"
  revoke_after = 1769689600000  # Unix timestamp in milliseconds for token auto-expiration
}
