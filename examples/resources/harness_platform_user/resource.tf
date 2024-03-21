terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

# Create user at account level
resource "harness_platform_user" "example" {
  email       = "niketan.test@harness.io"
  user_groups = ["PL_43763_v9"]
  role_bindings {
    resource_group_identifier = "_all_project_level_resources"
    role_identifier           = "_account_viewer"
    role_name                 = "Account Viewer"
    resource_group_name       = "All Account Level Resources Test"
    managed_role              = true
  }
}
