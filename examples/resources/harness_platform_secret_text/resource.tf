terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

resource "harness_platform_secret_text" "reference" {
  identifier  = "test_tf_secret_77"
  name        = "test_tf_secret_77"
  description = "test_tf_secret_description"
  tags        = ["foo:bar"]

  secret_manager_identifier = "GCP_SM"
  value_type                = "Reference"
  value                     = "sfdfsfewrewrxc"
  additional_metadata {
    values {
      version = "5"
    }
  }
}
