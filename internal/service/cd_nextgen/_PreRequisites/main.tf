terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

# Secret Variables
variable "github_token_value" {
  type = string
  sensitive = true
}

variable "harness_automation_github_token" {
  type = string
  sensitive = true
}

resource "harness_platform_secret_text" "TF_spot_account_id" {
  identifier                = "TF_spot_account_id"
  name                      = "TF_spot_account_id"
  description               = "This is a test Spot secret text"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"

  lifecycle {
    ignore_changes = [identifier]
  }
}

resource "harness_platform_secret_text" "TF_spot_api_token" {
  identifier                = "TF_spot_api_token"
  name                      = "TF_spot_api_token"
  description               = "This is a test Spot secret text"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"

  lifecycle {
    ignore_changes = [identifier]
  }
}

resource "harness_platform_secret_text" "TF_spot_api_token_ref" {
  identifier                = "TF_spot_api_token_ref"
  name                      = "TF_spot_api_token_ref"
  description               = "This is a test Spot secret text"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"

  lifecycle {
    ignore_changes = [identifier]
  }
}

resource "harness_platform_secret_text" "TF_Nexus_Password" {
  identifier                = "TF_Nexus_Password"
  name                      = "TF_Nexus_Password"
  description               = "This is a test secret text"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"

  lifecycle {
    ignore_changes = [identifier]
  }
}

resource "harness_platform_secret_text" "TF_git_bot_token" {
  identifier                = "TF_git_bot_token"
  name                      = "TF_git_bot_token"
  description               = "TF_git_bot_token"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = var.github_token_value

  lifecycle {
    ignore_changes = [identifier]
  }
}

resource "harness_platform_secret_text" "TF_harness_automation_github_token" {
  identifier                = "TF_harness_automation_github_token"
  name                      = "TF_harness_automation_github_token"
  description               = "TF_harness_automation_github_token"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = var.harness_automation_github_token

  lifecycle {
    ignore_changes = [identifier]
  }
}

resource "harness_platform_connector_github" "TF_GitX_connector" {
  identifier          = "TF_GitX_connector"
  name                = "TF_GitX_connector"
  description         = "TF_GitX_connector"
  tags                = ["ritek:test"]
  url                 = "https://github.com/harness-automation"
  connection_type     = "Account"
  validation_repo     = "pcf_practice"
  execute_on_delegate = false

  credentials {
    http {
      username  = "harness-automation"
      token_ref = "account.TF_harness_automation_github_token"
    }
  }

  api_authentication {
    token_ref = "account.TF_harness_automation_github_token"
  }

  lifecycle {
    ignore_changes = [identifier]
  }
}

resource "harness_platform_connector_github" "TF_open_repo_github_connector" {
  identifier          = "TF_open_repo_github_connector"
  name                = "TF_open_repo_github_connector"
  description         = "TF_open_repo_github_connector"
  tags                = ["ritek:test"]
  url                 = "https://github.com/harness-automation/open-repo"
  execute_on_delegate = false
  connection_type     = "Repo"

  credentials {
    http {
      username  = "admin"
      token_ref = "account.TF_harness_automation_github_token"
    }
  }

  api_authentication {
    token_ref = "account.TF_harness_automation_github_token"
  }

  lifecycle {
    ignore_changes = [identifier]
  }
}

resource "harness_platform_connector_github" "TF_Jajoo_github_connector" {
  identifier          = "TF_Jajoo_github_connector"
  name                = "TF_Jajoo_github_connector"
  description         = "TF_Jajoo_github_connector"
  tags                = ["ritek:test"]
  url                 = "https://github.com/harness-automation"
  connection_type     = "Account"
  validation_repo     = "jajoo_git.git"
  execute_on_delegate = false

  credentials {
    http {
      username  = "harness-automation"
      token_ref = "account.TF_harness_automation_github_token"
    }
  }

  api_authentication {
    token_ref = "account.TF_harness_automation_github_token"
  }

  lifecycle {
    ignore_changes = [identifier]
  }
}

resource "harness_platform_connector_github" "TF_TerraformResource_git_connector" {
  identifier          = "TF_TerraformResource_git_connector"
  name                = "TF_TerraformResource_git_connector"
  description         = "TF_TerraformResource_git_connector"
  tags                = ["ritek:test"]
  url                 = "https://github.com/harness-automation"
  connection_type     = "Account"
  validation_repo     = "pcf_practice"
  execute_on_delegate = false

  credentials {
    http {
      username  = "harness-automation"
      token_ref = "account.TF_harness_automation_github_token"
    }
  }

  api_authentication {
    token_ref = "account.TF_harness_automation_github_token"
  }

  lifecycle {
    ignore_changes = [identifier]
  }
}

resource "harness_platform_connector_github" "TF_github_account_level_delegate_connector" {
  identifier          = "TF_github_account_level_delegate_connector"
  name                = "TF_github_account_level_delegate_connector"
  description         = "TF_github_account_level_delegate_connector"
  tags                = ["ritek:test"]
  url                 = "https://github.com/harness-automation"
  connection_type     = "Account"
  validation_repo     = "Gitx-automation"
  execute_on_delegate = false

  credentials {
    http {
      username  = "harness-automation"
      token_ref = "account.TF_harness_automation_github_token"
    }
  }

  api_authentication {
    token_ref = "account.TF_harness_automation_github_token"
  }

  lifecycle {
    ignore_changes = [identifier]
  }
}

resource "harness_platform_connector_github" "TF_github_account_level_connector" {
  identifier          = "TF_github_account_level_connector"
  name                = "TF_github_account_level_connector"
  description         = "TF_github_account_level_connector"
  tags                = ["ritek:test"]
  url                 = "https://github.com/harness-automation"
  connection_type     = "Account"
  validation_repo     = "GitXTest3"
  execute_on_delegate = false

  credentials {
    http {
      username  = "harness-automation"
      token_ref = "account.TF_harness_automation_github_token"
    }
  }

  api_authentication {
    token_ref = "account.TF_harness_automation_github_token"
  }

  lifecycle {
    ignore_changes = [identifier]
  }
}
