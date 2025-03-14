resource "harness_platform_connector_gcp_secret_manager" "gcp_sm_manual" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  delegate_selectors = ["harness-delegate"]
  credentials_ref    = "account.${harness_platform_secret_text.test.id}"
}

resource "harness_platform_connector_gcp_secret_manager" "gcp_sm_inherit" {
  identifier            = "identifier"
  name                  = "name"
  description           = "test"
  tags                  = ["foo:bar"]
  delegate_selectors    = ["harness-delegate"]
  inherit_from_delegate = true
}

resource "harness_platform_connector_gcp_secret_manager" "gcp_sm_oidc_platform" {
  identifier          = "identifier"
  name                = "name"
  description         = "test"
  tags                = ["foo:bar"]
  execute_on_delegate = false
  oidc_authentication {
    workload_pool_id      = "harness-pool-test"
    provider_id           = "harness"
    gcp_project_id        = "1234567"
    service_account_email = "harness.sample@iam.gserviceaccount.com"
  }
}

resource "harness_platform_connector_gcp_secret_manager" "gcp_sm_oidc_delegate" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  default     = true

  delegate_selectors = ["harness-delegate"]
  oidc_authentication {
    workload_pool_id      = "harness-pool-test"
    provider_id           = "harness"
    gcp_project_id        = "1234567"
    service_account_email = "harness.sample@iam.gserviceaccount.com"
  }
}
