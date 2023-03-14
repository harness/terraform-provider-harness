# Create Terraform Cloud connector using API token as secret

resource "harness_platform_connector_terraform_cloud" "terraform_cloud" {
  identifier         = "example_terraform_cloud_connector"
  name               = "Example terraform cloud connector"
  description        = "description of terraform cloud connector"
  tags               = ["foo:bar"]
  delegate_selectors = ["harness-delegate"]
  credentials {
    api_token {
      api_token_ref = "account.TEST_terraform_cloud_api_token"
    }
  }
}

# Add connectivity mode by providing execute_on_delegate value. Default is to execute on Delegate

resource "harness_platform_connector_terraform_cloud" "terraform_cloud" {
  identifier          = "example_terraform_cloud_connector"
  name                = "Example terraform cloud connector"
  description         = "description of terraform cloud connector"
  delegate_selectors  = ["harness-delegate"]
  tags                = ["foo:bar"]
  execute_on_delegate = false
  credentials {
    api_token {
      api_token_ref = "account.TEST_terraform_cloud_api_token"
    }
  }
}
