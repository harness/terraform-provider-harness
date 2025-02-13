resource "harness_platform_infra_module" "example" {
  identifier  = "example"
  name        = "example"
  org_id      = harness_platform_organization.test.id
  project_id  = harness_platform_project.test.id
  description = "some description"

  environment_variable {
    key        = "key1"
    value      = "value1"
    value_type = "string"
  }
  environment_variable {
    key        = "key2"
    value      = "harness_platform_secret_text.test.id"
    value_type = "secret"
  }

  terraform_variable {
    key        = "key1"
    value      = "1111"
    value_type = "string"
  }
  terraform_variable {
    key        = "key2"
    value      = "1111u"
    value_type = "string"
  }

  terraform_variable_file {
    repository           = "https://github.com/org/repo"
    repository_branch    = "main"
    repository_path      = "tf/aws/basic"
    repository_connector = "harness_platform_connector_github.test.id"
  }
  terraform_variable_file {
    repository           = "https://github.com/org/repo"
    repository_branch    = "br2"
    repository_path      = "tf/aws/basic"
    repository_connector = "harness_platform_connector_github.test.id"
  }

  connector {
    connector_ref = "harness_platform_connector_aws.test.id"
    type          = "aws"
  }
  connector {
    connector_ref = "harness_platform_connector_azure.test.id"
    type          = "azure"
  }

}
