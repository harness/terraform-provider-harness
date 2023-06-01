# Create token for account level apikey
resource "harness_platform_token" "test" {
  identifier  = "test_token"
  name        = "test token"
  parent_id   = "apikey_parent_id"
  account_id  = "account_id"
  apikey_type = "USER"
  apikey_id   = "apikey_id"
}

# Create token for org level apikey
resource "harness_platform_token" "test" {
  identifier  = "test_token"
  name        = "test token"
  parent_id   = "apikey_parent_id"
  account_id  = "account_id"
  org_id      = "org_id"
  apikey_type = "USER"
  apikey_id   = "apikey_id"
}

# Create token for project level apikey
resource "harness_platform_token" "test" {
  identifier  = "test_token"
  name        = "test token"
  parent_id   = "apikey_parent_id"
  account_id  = "account_id"
  org_id      = "org_id"
  project_id  = "project_id"
  apikey_type = "USER"
  apikey_id   = "apikey_id"
}