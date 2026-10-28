# Token auth — direct Anthropic API key
resource "harness_platform_connector_anthropic_model" "token" {
  identifier  = "example_anthropic_token"
  name        = "example-anthropic-token"
  description = "Anthropic connector with Token auth"
  tags        = ["env:dev"]

  url                    = "https://api.anthropic.com"
  model                  = "claude-opus-4-6"
  delegate_selectors     = ["harness-delegate"]
  execute_on_delegate    = true
  ignore_test_connection = false

  auth {
    auth_type = "Token"
    token {
      token_ref = "account.anthropic_api_key"
    }
  }
}

# BedrockApiKey auth — AWS Bedrock with API key + region
resource "harness_platform_connector_anthropic_model" "bedrock" {
  identifier  = "example_anthropic_bedrock"
  name        = "example-anthropic-bedrock"
  description = "Anthropic connector with BedrockApiKey auth"
  tags        = ["env:dev"]

  url                    = "https://api.anthropic.com"
  model                  = "claude-opus-4-6"
  delegate_selectors     = ["harness-delegate"]
  execute_on_delegate    = true
  ignore_test_connection = true

  auth {
    auth_type = "BedrockApiKey"
    bedrock_api_key {
      api_key_ref = "account.bedrock_api_key"
      region      = "us-east-1"
    }
  }
}

# CloudProvider auth — delegate to an existing AWS connector
resource "harness_platform_connector_anthropic_model" "cloud_provider" {
  identifier  = "example_anthropic_cloud_provider"
  name        = "example-anthropic-cloud-provider"
  description = "Anthropic connector with AWS CloudProvider auth"
  tags        = ["env:dev"]

  url                    = "https://api.anthropic.com"
  model                  = "claude-opus-4-6"
  delegate_selectors     = ["harness-delegate"]
  execute_on_delegate    = true
  ignore_test_connection = true

  auth {
    auth_type = "CloudProvider"
    cloud_provider {
      type          = "AWS"
      connector_ref = "account.aws_connector"
    }
  }
}
