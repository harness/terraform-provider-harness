# Create a Terraform Cloud connector by using an API token as a secret.

resource "harness_platform_connector_terraform_cloud" "terraform_cloud" {
  identifier         = "example_terraform_cloud_connector"
  name               = "Example terraform cloud connector"
  description        = "description of terraform cloud connector"
  tags               = ["foo:bar"]
  delegate_selectors = ["harness-delegate"]
  url                = "https://app.terraform.io/"
  credentials {
    api_token {
      api_token_ref = "account.TEST_terraform_cloud_api_token"
    }
  }
}

# Specify the connectivity mode by setting execute_on_delegate to true or false. The default mode executes on the delegate.

resource "harness_platform_connector_terraform_cloud" "terraform_cloud" {
  identifier          = "example_terraform_cloud_connector"
  name                = "Example terraform cloud connector"
  description         = "description of terraform cloud connector"
  delegate_selectors  = ["harness-delegate"]
  tags                = ["foo:bar"]
  execute_on_delegate = false
  url                = "https://app.terraform.io/"
  credentials {
    api_token {
      api_token_ref = "account.TEST_terraform_cloud_api_token"
    }
  }
}
