# Example 1: Username/Password Authentication (New Block Format)
resource "harness_platform_connector_splunk" "username_password" {
  identifier  = "splunk_userpass"
  name        = "Splunk Username/Password"
  description = "Splunk connector with username/password authentication"
  tags        = ["env:production"]

  url                = "https://splunk.company.com:8089"
  delegate_selectors = ["harness-delegate"]
  account_id         = "splunk_account_id"
  
  username_password {
    username     = "splunk_user"
    password_ref = "account.splunk_password"
  }
}

# Example 2: Bearer Token Authentication
resource "harness_platform_connector_splunk" "bearer_token" {
  identifier  = "splunk_bearer"
  name        = "Splunk Bearer Token"
  description = "Splunk connector with bearer token authentication"
  tags        = ["env:production"]

  url                = "https://splunk.company.com:8089"
  delegate_selectors = ["harness-delegate"]
  account_id         = "splunk_account_id"
  
  bearer_token {
    bearer_token_ref = "account.splunk_bearer_token"
  }
}

# Example 3: HEC Token Authentication
resource "harness_platform_connector_splunk" "hec_token" {
  identifier  = "splunk_hec"
  name        = "Splunk HEC Token"
  description = "Splunk connector with HEC token authentication"
  tags        = ["env:production"]

  url                = "https://splunk.company.com:8088"
  delegate_selectors = ["harness-delegate"]
  account_id         = "splunk_account_id"
  
  hec_token {
    hec_token_ref = "account.splunk_hec_token"
  }
}

# Example 4: No Authentication
resource "harness_platform_connector_splunk" "no_auth" {
  identifier  = "splunk_no_auth"
  name        = "Splunk No Auth"
  description = "Splunk connector without authentication"
  tags        = ["env:development"]

  url                = "https://splunk-dev.company.com:8089"
  delegate_selectors = ["harness-delegate"]
  account_id         = "splunk_account_id"
  
  no_authentication {}
}

# Example 5: Legacy Format (Deprecated but still supported)
resource "harness_platform_connector_splunk" "legacy" {
  identifier  = "splunk_legacy"
  name        = "Splunk Legacy"
  description = "Splunk connector using deprecated flat schema"
  tags        = ["deprecated"]

  url                = "https://splunk.company.com:8089"
  delegate_selectors = ["harness-delegate"]
  account_id         = "splunk_account_id"
  username           = "username"           # Deprecated
  password_ref       = "account.secret_id"  # Deprecated
}
