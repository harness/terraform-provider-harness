# Credential manual
resource "harness_platform_connector_gcp" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  manual {
    secret_key_ref     = "account.secret_id"
    delegate_selectors = ["harness-delegate"]
  }
}

# Credentials inherit_from_delegate
resource "harness_platform_connector_gcp" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  inherit_from_delegate {
    delegate_selectors = ["harness-delegate"]
  }
}

# Create Gcp connector using Oidc Authentication

resource "harness_platform_connector_gcp" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  oidc_authentication {
    workload_pool_id = "harness-pool-test"
    provider_id = "harness"
    gcp_project_id = "1234567"
    service_account_email = "harness.sample.iam.gserviceaccount.com"
    delegate_selectors = ["harness-delegate"]
  }
}
