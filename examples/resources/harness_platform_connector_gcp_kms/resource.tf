terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

resource "harness_platform_secret_text" "test" {
  identifier                = "%[1]s"
  name                      = "%[2]s"
  description               = "test"
  tags                      = ["foo:bar"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Reference"
  value                     = "secret"
}

resource "harness_platform_connector_gcp_kms" "gcp_kms_manual" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  region         = "us-west1"
  gcp_project_id = "1234567"
  key_ring       = "key_ring"
  key_name       = "key_name"

  manual {
    credentials        = "account.${harness_platform_secret_text.test.id}"
    delegate_selectors = ["harness-delegate"]
  }
}

resource "harness_platform_connector_gcp_kms" "gcp_kms_oidc_platform" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  region         = "us-west1"
  gcp_project_id = "1234567"
  key_ring       = "key_ring"
  key_name       = "key_name"

  execute_on_delegate = false

  oidc_authentication {
    workload_pool_id      = "harness-pool-test"
    provider_id           = "harness"
    gcp_project_id        = "1234567"
    service_account_email = "harness.sample@iam.gserviceaccount.com"
  }
}

resource "harness_platform_connector_gcp_kms" "gcp_kms_oidc_delegate" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  region         = "us-west1"
  gcp_project_id = "1234567"
  key_ring       = "key_ring"
  key_name       = "key_name"

  oidc_authentication {
    workload_pool_id      = "harness-pool-test"
    provider_id           = "harness"
    gcp_project_id        = "1234567"
    service_account_email = "harness.sample@iam.gserviceaccount.com"
    delegate_selectors    = ["harness-delegate"]
  }
}

resource "harness_platform_connector_gcp_kms" "gcp_kms_oidc_delegate_default" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  region         = "us-west1"
  gcp_project_id = "1234567"
  key_ring       = "key_ring"
  key_name       = "key_name"

  default = true

  oidc_authentication {
    workload_pool_id      = "harness-pool-test"
    provider_id           = "harness"
    gcp_project_id        = "1234567"
    service_account_email = "harness.sample@iam.gserviceaccount.com"
    delegate_selectors    = ["harness-delegate"]
  }
}
