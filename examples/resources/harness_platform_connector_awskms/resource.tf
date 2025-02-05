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

# Credentials manual as Default Secret Manager
resource "harness_platform_connector_awskms" "test" {
  identifier  = "identifer"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]


  arn_ref            = "account.secret_id"
  region             = "us-east-1"
  delegate_selectors = ["harness-delegate"]
  default            = true
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

# Credentials OIDC using Harness Platform
resource "harness_platform_connector_awskms" "test" {
  identifier          = "%[1]s"
  name                = "%[1]s"
  description         = "test"
  tags                = ["foo:bar"]
  arn_ref             = "account.secret_id"
  region              = "us-east-1"
  execute_on_delegate = false
  credentials {
    oidc_authentication {
      iam_role_arn = "somerolearn"
    }
  }
}

# Credentials OIDC using Delegate
resource "harness_platform_connector_awskms" "test" {
  identifier         = "%[1]s"
  name               = "%[1]s"
  description        = "test"
  tags               = ["foo:bar"]
  arn_ref            = "account.secret_id"
  region             = "us-east-1"
  delegate_selectors = ["harness-delegate"]
  credentials {
    oidc_authentication {
      iam_role_arn = "somerolearn"
    }
  }
}
