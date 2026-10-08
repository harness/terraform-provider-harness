# Token authentication
resource "harness_platform_connector_openai_model" "token" {
  identifier             = "example_openai_token"
  name                   = "example-openai-token"
  description            = "OpenAI connector with Token auth"
  url                    = "https://api.openai.com"
  model                  = "gpt-4o"
  execute_on_delegate    = true
  ignore_test_connection = false

  auth {
    auth_type = "Token"
    token {
      token_ref = "account.openai_api_key"
    }
  }
}

# Vertex AI authentication
resource "harness_platform_connector_openai_model" "vertex" {
  identifier             = "example_openai_vertex"
  name                   = "example-openai-vertex"
  description            = "OpenAI connector with Vertex AI auth"
  url                    = "https://api.openai.com"
  model                  = "gpt-4o"
  execute_on_delegate    = true
  ignore_test_connection = true
  delegate_selectors     = ["mydelegate"]

  auth {
    auth_type = "Vertex"
    vertex {
      service_account_key_ref = "account.gcp_sa_key"
      project_id              = "my-gcp-project"
      region                  = "us-central1"
    }
  }
}
