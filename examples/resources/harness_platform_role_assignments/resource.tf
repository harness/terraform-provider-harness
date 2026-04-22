//To create a role binding in service account
resource "harness_platform_role_assignments" "example1" {
  org_id                    = "org_id"
  project_id                = "project_id"
  resource_group_identifier = "_all_project_level_resources"
  role_identifier           = "_project_viewer"
  principal {
    identifier = harness_platform_service_account.test.id
    type       = "SERVICE_ACCOUNT"
  }
  disabled = false
  managed  = false
}

//To create a role binding in user group
resource "harness_platform_role_assignments" "example1" {
  org_id                    = "org_id"
  project_id                = "project_id"
  resource_group_identifier = "_all_project_level_resources"
  role_identifier           = "_project_viewer"
  principal {
    identifier = harness_platform_usergroup.test.id
    type       = "USER_GROUP"
  }
  disabled = false
  managed  = false
}

resource "harness_platform_role_assignments" "example2" {
  identifier                = "identifier"
  org_id                    = "org_id"
  project_id                = "project_id"
  resource_group_identifier = "_all_project_level_resources"
  role_identifier           = "_project_viewer"
  principal {
    identifier = "user_id"
    type       = "USER"
  }
  disabled = false
  managed  = false
}

resource "harness_platform_role_assignments" "example2" {
  org_id                    = "org_id"
  project_id                = "project_id"
  resource_group_identifier = "_all_project_level_resources"
  role_identifier           = "_project_viewer"
  principal {
    identifier = "service_id"
    type       = "SERVICE"
  }
  disabled = false
  managed  = false
}

resource "harness_platform_role_assignments" "example2" {
  org_id                    = "org_id"
  project_id                = "project_id"
  resource_group_identifier = "_all_project_level_resources"
  role_identifier           = "_project_viewer"
  principal {
    identifier = "api_key_id"
    type       = "API_KEY"
  }
  disabled = false
  managed  = false
}

//To create a role binding using role_reference (e.g. an org-level role assigned at project scope)
resource "harness_platform_role_assignments" "example3" {
  identifier                = "identifier"
  org_id                    = "org_id"
  project_id                = "project_id"
  resource_group_identifier = "_all_project_level_resources"
  role_identifier           = "org_role_identifier"
  role_reference {
    identifier  = "org_role_identifier"
    scope_level = "organization"
  }
  principal {
    identifier = harness_platform_service_account.test.id
    type       = "SERVICE_ACCOUNT"
  }
  disabled = false
  managed  = false
}
