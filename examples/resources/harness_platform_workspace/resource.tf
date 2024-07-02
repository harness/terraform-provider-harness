resource "harness_platform_workspace" "example" {
  name                    = "example"
  identifier              = "example"
  org_id                  = harness_platform_organization.test.id
  project_id              = harness_platform_project.test.id
  provisioner_type        = "terraform"
  provisioner_version     = "1.5.6"
  repository              = "https://github.com/org/repo"
  repository_branch       = "main"
  repository_path         = "tf/aws/basic"
  cost_estimation_enabled = true
  provider_connector      = harness_platform_connector_github.test.id
  repository_connector    = harness_platform_connector_github.test.id

  terraform_variable {
    key        = "key1"
    value      = "val1"
    value_type = "string"
  }
  terraform_variable {
    key        = "key2"
    value      = "val2"
    value_type = "string"
  }

  environment_variable {
    key        = "key1"
    value      = "val1"
    value_type = "string"
  }
  environment_variable {
    key        = "key2"
    value      = "val2"
    value_type = "string"
  }

  terraform_variable_file {
    repository           = "https://github.com/org/repo"
    repository_branch    = "main"
    repository_path      = "tf/gcp/terraform.tfvars"
    repository_connector = harness_platform_connector_github.test.id
  }
  terraform_variable_file {
    repository           = "https://github.com/org/repo"
    repository_commit    = "349d90bb9c90f4a3482981c259080de31609e6f6"
    repository_path      = "tf/aws/terraform.tfvars"
    repository_connector = harness_platform_connector_github.test.id
  }
}
