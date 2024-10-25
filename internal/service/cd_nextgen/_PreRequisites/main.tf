terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

variable "github_token_value" {
  type = string
}

resource "harness_platform_secret_text" "TEST_spot_account_id" {
  identifier                = "TEST_spot_account_id"
  name                      = "TEST_spot_account_id"
  description               = "This is a test Spot secret text"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"
}

resource "harness_platform_secret_text" "TEST_spot_api_token" {
  identifier                = "TEST_spot_api_token"
  name                      = "TEST_spot_api_token"
  description               = "This is a test Spot secret text"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"
  depends_on                = [harness_platform_secret_text.TEST_spot_account_id]
}

resource "harness_platform_secret_text" "TEST_api_token_ref" {
  identifier                = "TEST_api_token_ref"
  name                      = "TEST_api_token_ref"
  description               = "This is a test Spot secret text"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"
  depends_on                = [harness_platform_secret_text.TEST_spot_api_token]
}

resource "harness_platform_secret_text" "doNotDeleteHSM" {
  identifier                = "doNotDeleteHSM"
  name                      = "doNotDeleteHSM"
  description               = "This is a test secret text"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"
  depends_on                = [harness_platform_secret_text.TEST_api_token_ref]
}

resource "harness_platform_secret_text" "gitbotharnesstoken" {
  identifier                = "gitbotharnesstoken"
  name                      = "gitbotharnesstoken"
  description               = "gitbotharnesstoken"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = var.github_token_value
  depends_on                = [harness_platform_secret_text.doNotDeleteHSM]
}

resource "harness_platform_connector_github" "DoNotDeleteGitX" {
  identifier  = "DoNotDeleteGitX"
  name        = "DoNotDeleteGitX"
  description = "DoNotDeleteGitX"
  tags        = ["ritek:test"]

  url             = "https://github.com/sourabh-awashti/pcf_practice"
  connection_type = "Repo"
  credentials {
    http {
      anonymous {}
    }
  }
  depends_on = [harness_platform_secret_text.gitbotharnesstoken]
}

resource "harness_platform_connector_github" "Jajoo" {
  identifier  = "Jajoo"
  name        = "Jajoo"
  description = "Jajoo"
  tags        = ["ritek:test"]

  url             = "https://github.com/wings-software/jajoo_git"
  connection_type = "Repo"
  credentials {
    http {
      username  = "admin"
      token_ref = "account.gitbotharnesstoken"
    }
  }
  depends_on = [harness_platform_connector_github.DoNotDeleteGitX]
}

resource "harness_platform_connector_git" "DoNotDeleteRTerraformResource" {
  identifier       = "DoNotDeleteRTerraformResource"
  name             = "DoNotDeleteRTerraformResource"
  description      = "DoNotDeleteRTerraformResource"
  tags             = ["ritek:test"]

  url              = "https://github.com/wings-software/terraform-test"
  connection_type  = "Repo"
  credentials {
    http {
      username     = "admin"
      password_ref = "account.gitbotharnesstoken"
    }
  }
  depends_on = [harness_platform_connector_github.Jajoo]
}
