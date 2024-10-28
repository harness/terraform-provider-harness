# main.tf

terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

variable "TF_VAR_github_token_value" {
  type = string
}

variable "TF_VAR_harness_automation_github_token" {
  type = string
}

resource "harness_platform_project" "TF_Pipeline_Test" {
		identifier = "TF_Pipeline_Test"
		name = "TF_Pipeline_Test"
		color = "#0063F7"
		org_id = "default"
}

resource "harness_platform_secret_text" "TF_spot_account_id" {
  identifier                = "TF_spot_account_id"
  name                      = "TF_spot_account_id"
  description               = "This is a test Spot secret text"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"
  depends_on                = [harness_platform_project.TF_Pipeline_Test]
}

resource "harness_platform_secret_text" "TF_spot_api_token" {
  identifier                = "TF_spot_api_token"
  name                      = "TF_spot_api_token"
  description               = "This is a test Spot secret text"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"
  depends_on                = [harness_platform_secret_text.TF_spot_account_id]
}

resource "harness_platform_secret_text" "TF_spot_api_token_ref" {
  identifier                = "TF_spot_api_token_ref"
  name                      = "TF_spot_api_token_ref"
  description               = "This is a test Spot secret text"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"
  depends_on                = [harness_platform_secret_text.TF_spot_api_token]
}

resource "harness_platform_secret_text" "TF_Nexus_Password" {
  identifier                = "TF_Nexus_Password"
  name                      = "TF_Nexus_Password"
  description               = "This is a test secret text"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "my_secret_value"
  depends_on                = [harness_platform_secret_text.TF_spot_api_token_ref]
}

resource "harness_platform_secret_text" "TF_git_bot_token" {
  identifier                = "TF_git_bot_token"
  name                      = "TF_git_bot_token"
  description               = "TF_git_bot_token"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = var.TF_VAR_github_token_value
  depends_on                = [harness_platform_secret_text.TF_Nexus_Password]
}

resource "harness_platform_secret_text" "TF_harness_automation_github_token" {
  identifier                = "TF_harness_automation_github_token"
  name                      = "TF_harness_automation_github_token"
  description               = "TF_harness_automation_github_token"
  tags                      = ["ritek:test"]
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = var.TF_VAR_harness_automation_github_token
  depends_on                = [harness_platform_secret_text.TF_git_bot_token]
}

resource "harness_platform_connector_github" "TF_GitX_connector" {
  identifier  = "TF_GitX_connector"
  name        = "TF_GitX_connector"
  description = "TF_GitX_connector"
  tags        = ["ritek:test"]

  url             = "https://github.com/sourabh-awashti/pcf_practice"
  connection_type = "Repo"
  execute_on_delegate = false
  credentials {
    http {
      anonymous {}
    }
  }
  depends_on = [harness_platform_secret_text.TF_harness_automation_github_token]
}

resource "harness_platform_connector_github" "TF_open_repo_github_connector" {
  identifier  = "TF_open_repo_github_connector"
  name        = "TF_open_repo_github_connector"
  description = "TF_open_repo_github_connector"
  tags        = ["ritek:test"]

  url             = "https://github.com/vtxorxwitty/open-repo"
  execute_on_delegate = false
  connection_type = "Repo"
    credentials {
      http {
        username  = "admin"
        token_ref = "account.TF_git_bot_token"
      }
    }
    api_authentication {
        token_ref = "account.TF_git_bot_token"
    }
  depends_on = [harness_platform_connector_github.TF_GitX_connector]
}

resource "harness_platform_connector_github" "TF_Jajoo_github_connector" {
  identifier  = "TF_Jajoo_github_connector"
  name        = "TF_Jajoo_github_connector"
  description = "TF_Jajoo_github_connector"
  tags        = ["ritek:test"]

  url             = "https://github.com/wings-software/jajoo_git"
  connection_type = "Repo"
  credentials {
    http {
      username  = "admin"
      token_ref = "account.TF_git_bot_token"
    }
  }
  api_authentication {
      token_ref = "account.TF_git_bot_token"
  }
  depends_on = [harness_platform_connector_github.TF_open_repo_github_connector]
}

resource "harness_platform_connector_git" "TF_TerraformResource_git_connector" {
  identifier       = "TF_TerraformResource_git_connector"
  name             = "TF_TerraformResource_git_connector"
  description      = "TF_TerraformResource_git_connector"
  tags             = ["ritek:test"]

  url              = "https://github.com/wings-software/terraform-test"
  connection_type  = "Repo"
  credentials {
    http {
      username     = "admin"
      password_ref = "account.TF_git_bot_token"
    }
  }
  depends_on = [harness_platform_connector_github.TF_Jajoo_github_connector]
}

resource "harness_platform_connector_github" "TF_github_account_level_delegate_connector" {
  identifier  = "TF_github_account_level_delegate_connector"
  name        = "TF_github_account_level_delegate_connector"
  description = "TF_github_account_level_delegate_connector"
  tags        = ["ritek:test"]

  url             = "https://github.com/harness-automation"
  connection_type = "Account"
  validation_repo = "Gitx-automation"
  credentials {
    http {
      username  = "harness-automation"
      token_ref = "account.TF_harness_automation_github_token"
    }
  }
  api_authentication {
      token_ref = "account.TF_harness_automation_github_token"
  }
  depends_on = [harness_platform_connector_git.TF_TerraformResource_git_connector]
}

resource "harness_platform_connector_github" "TF_github_account_level_connector" {
  identifier  = "TF_github_account_level_connector"
  name        = "TF_github_account_level_connector"
  description = "TF_github_account_level_connector"
  tags        = ["ritek:test"]

  url             = "https://github.com/harness-automation"
  connection_type = "Account"
  validation_repo = "GitXTest3"
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
  depends_on = [harness_platform_connector_github.TF_github_account_level_delegate_connector]
}
