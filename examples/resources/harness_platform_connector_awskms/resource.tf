# Credentials assume_role
resource "harness_platform_connector_awskms" "test" {
  identifier  = "identifer"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  arn_ref            = "account.secret_id"
  region             = "us-east-1"
  delegate_selectors = ["harness-delegate"]
  credentials {
    assume_role {
      role_arn    = "somerolearn"
      external_id = "externalid"
      duration    = 900
    }
  }
}

# Credentials manual
resource "harness_platform_connector_awskms" "test" {
  identifier  = "identifer"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]


  arn_ref            = "account.secret_id"
  region             = "us-east-1"
  delegate_selectors = ["harness-delegate"]
  credentials {
    manual {
      secret_key_ref = "account.secret_id"
      access_key_ref = "account.secret_id"
    }
  }
}

# Credentials inherit_from_delegate
resource "harness_platform_connector_awskms" "test" {
  identifier  = "identifer"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  arn_ref            = "account.secret_id"
  region             = "us-east-1"
  delegate_selectors = ["harness-delegate"]
  credentials {
    inherit_from_delegate = true
  }
}
