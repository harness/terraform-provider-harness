# Create Spot connector using permanent token and spot account Id as plain text

resource "harness_platform_connector_spot" "spot" {
  identifier  = "example_spot_cloud_provider"
  name        = "Example spot cloud provider"
  description = "description of spot connector"
  tags        = ["foo:bar"]
  permanent_token {
    spot_account_id    = "<my-account-id>"
    api_token_ref      = "account.TEST_spot_api_token"
    delegate_selectors = ["harness-delegate"]
  }
}


# Create Spot connector using permanent token and spot account Id as secret

resource "harness_platform_connector_spot" "spot" {
  identifier  = "example_spot_cloud_provider"
  name        = "Example spot cloud provider"
  description = "description of spot connector"
  tags        = ["foo:bar"]
  permanent_token {
    spot_account_id_ref = "account.TEST_spot_account_id"
    api_token_ref       = "account.TEST_spot_api_token"
    delegate_selectors  = ["harness-delegate"]
  }
}

# Add connectivity mode by providing execute_on_delegate value. Default is to execute on Delegate

resource "harness_platform_connector_spot" "spot" {
  identifier  = "example_spot_cloud_provider"
  name        = "Example spot cloud provider"
  description = "description of spot connector"
  tags        = ["foo:bar"]
  permanent_token {
    spot_account_id     = "<my-account-id>"
    api_token_ref       = "account.TEST_spot_api_token"
    delegate_selectors  = ["harness-delegate"]
    execute_on_delegate = false
  }
}
