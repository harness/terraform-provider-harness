# main.tf

terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

resource "harness_platform_secret_text" "TEST_spot_account_id" {
  identifier                = "TEST_spot_account_id"
  name                      = "TEST_spot_account_id"
  description               = "This is a test Spot secret text"
  tags                      = ["foo:bar"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"
}
resource "harness_platform_secret_text" "TEST_spot_api_token" {
  identifier                = "TEST_spot_api_token"
  name                      = "TEST_spot_api_token"
  description               = "This is a test Spot secret text"
  tags                      = ["foo:bar"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"
}

resource "harness_platform_secret_text" "TEST_api_token_ref" {
  identifier                = "TEST_api_token_ref"
  name                      = "TEST_api_token_ref"
  description               = "This is a test Spot secret text"
  tags                      = ["foo:bar"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"
}

resource "harness_platform_secret_text" "doNotDeleteHSM" {
  identifier                = "doNotDeleteHSM"
  name                      = "doNotDeleteHSM"
  description               = "This is a test secret text"
  tags                      = ["foo:bar"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"
}

# Additional resources can be added as needed
