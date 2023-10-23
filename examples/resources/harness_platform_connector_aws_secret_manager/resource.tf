# Credentials inherit_from_delegate
resource "harness_platform_connector_aws_secret_manager" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  secret_name_prefix = "test"
  region             = "us-east-1"
  delegate_selectors = ["harness-delegate"]
  credentials {
    inherit_from_delegate = true
  }
}

# Credentials manual
resource "harness_platform_connector_aws_secret_manager" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  secret_name_prefix = "test"
  region             = "us-east-1"
  delegate_selectors = ["harness-delegate"]
  credentials {
    manual {
      secret_key_ref = "account.secret_id"
      access_key_ref = "account.secret_id"
    }
  }
}

# Credentials assume_role
resource "harness_platform_connector_aws_secret_manager" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  secret_name_prefix = "test"
  region             = "us-east-1"
  delegate_selectors = ["harness-delegate"]
  default            = true
  credentials {
    assume_role {
      role_arn    = "somerolearn"
      external_id = "externalid"
      duration    = 900
    }
  }
}
