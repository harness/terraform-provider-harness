terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

resource "harness_platform_secret_text" "example" {
    identifier                = "test_gcp_ref_sec_tf"
    name                      = "test_gcp_ref_sec_tf"
    secret_manager_identifier = "account.GCP_SM"
    tags                      = []
    value                     = "(sensitive value)"
    value_type                = "Reference"

    additional_metadata {
        values {
            version = "1"
        }
    }
}
