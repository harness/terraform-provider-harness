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

# Data sources with try() function to safely handle non-existent resources
locals {
  spot_account_id_exists                  = try(data.harness_platform_secret_text.existing_spot_account_id.id, null) != null
  spot_api_token_exists                   = try(data.harness_platform_secret_text.existing_spot_api_token.id, null) != null
  spot_api_token_ref_exists               = try(data.harness_platform_secret_text.existing_spot_api_token_ref.id, null) != null
  nexus_password_exists                   = try(data.harness_platform_secret_text.existing_nexus_password.id, null) != null
  git_bot_token_exists                    = try(data.harness_platform_secret_text.existing_git_bot_token.id, null) != null
  harness_automation_github_token_exists  = try(data.harness_platform_secret_text.existing_harness_automation_github_token.id, null) != null
  
  gitx_connector_exists                    = try(data.harness_platform_connector_github.existing_gitx_connector.id, null) != null
  open_repo_connector_exists               = try(data.harness_platform_connector_github.existing_open_repo_connector.id, null) != null
  jajoo_connector_exists                   = try(data.harness_platform_connector_github.existing_jajoo_connector.id, null) != null
  terraform_resource_connector_exists      = try(data.harness_platform_connector_github.existing_terraform_resource_connector.id, null) != null
  account_level_delegate_connector_exists  = try(data.harness_platform_connector_github.existing_account_level_delegate_connector.id, null) != null
  account_level_connector_exists           = try(data.harness_platform_connector_github.existing_account_level_connector.id, null) != null
}

# Data sources to check if resources exist
data "harness_platform_secret_text" "existing_spot_account_id" {
  identifier = "TF_spot_account_id"
  name       = "TF_spot_account_id"
}

data "harness_platform_secret_text" "existing_spot_api_token" {
  identifier = "TF_spot_api_token"
  name       = "TF_spot_api_token"
}

data "harness_platform_secret_text" "existing_spot_api_token_ref" {
  identifier = "TF_spot_api_token_ref"
  name       = "TF_spot_api_token_ref"
}

data "harness_platform_secret_text" "existing_nexus_password" {
  identifier = "TF_Nexus_Password"
  name       = "TF_Nexus_Password"
}

data "harness_platform_secret_text" "existing_git_bot_token" {
  identifier = "TF_git_bot_token"
  name       = "TF_git_bot_token"
}

data "harness_platform_secret_text" "existing_harness_automation_github_token" {
  identifier = "TF_harness_automation_github_token"
  name       = "TF_harness_automation_github_token"
}

resource "harness_platform_secret_text" "TF_spot_account_id" {
  count                     = local.spot_account_id_exists ? 0 : 1
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
  count                     = local.spot_api_token_exists ? 0 : 1
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
  count                     = local.spot_api_token_ref_exists ? 0 : 1
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
  count                     = local.nexus_password_exists ? 0 : 1
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
  count                     = local.git_bot_token_exists ? 0 : 1
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
  count                     = local.harness_automation_github_token_exists ? 0 : 1
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

# Data sources to check if connectors exist
data "harness_platform_connector_github" "existing_gitx_connector" {
  identifier = "TF_GitX_connector"
  name       = "TF_GitX_connector"
}

data "harness_platform_connector_github" "existing_open_repo_connector" {
  identifier = "TF_open_repo_github_connector"
  name       = "TF_open_repo_github_connector"
}

data "harness_platform_connector_github" "existing_jajoo_connector" {
  identifier = "TF_Jajoo_github_connector"
  name       = "TF_Jajoo_github_connector"
}

data "harness_platform_connector_github" "existing_terraform_resource_connector" {
  identifier = "TF_TerraformResource_git_connector"
  name       = "TF_TerraformResource_git_connector"
}

data "harness_platform_connector_github" "existing_account_level_delegate_connector" {
  identifier = "TF_github_account_level_delegate_connector"
  name       = "TF_github_account_level_delegate_connector"
}

data "harness_platform_connector_github" "existing_account_level_connector" {
  identifier = "TF_github_account_level_connector"
  name       = "TF_github_account_level_connector"
}

# Modified connector resources with improved conditionals
resource "harness_platform_connector_github" "TF_GitX_connector" {
  count = !local.gitx_connector_exists && local.harness_automation_github_token_exists ? 1 : 0

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
  count = !local.open_repo_connector_exists && local.harness_automation_github_token_exists ? 1 : 0

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
  count = !local.jajoo_connector_exists && local.git_bot_token_exists ? 1 : 0

  identifier          = "TF_Jajoo_github_connector"
  name                = "TF_Jajoo_github_connector"
  description         = "TF_Jajoo_github_connector"
  tags                = ["ritek:test"]
  url                 = "https://github.com/wings-software/jajoo_git"
  connection_type     = "Repo"

  credentials {
    http {
      username  = "admin"
      token_ref = "account.TF_git_bot_token"
    }
  }

  api_authentication {
    token_ref = "account.TF_git_bot_token"
  }

  lifecycle {
    ignore_changes = [identifier]
  }
}

resource "harness_platform_connector_github" "TF_TerraformResource_git_connector" {
  count = !local.terraform_resource_connector_exists && local.harness_automation_github_token_exists ? 1 : 0

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
  count = !local.account_level_delegate_connector_exists && local.harness_automation_github_token_exists ? 1 : 0

  identifier          = "TF_github_account_level_delegate_connector"
  name                = "TF_github_account_level_delegate_connector"
  description         = "TF_github_account_level_delegate_connector"
  tags                = ["ritek:test"]
  url                 = "https://github.com/harness-automation"
  connection_type     = "Account"
  validation_repo     = "Gitx-automation"

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
  count = !local.account_level_connector_exists && local.harness_automation_github_token_exists ? 1 : 0

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
