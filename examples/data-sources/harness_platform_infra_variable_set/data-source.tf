# Look up an account level Variable Set. Omit both org_id and project_id.
data "harness_platform_infra_variable_set" "account_level" {
  identifier = "account_variable_set"
}

# Look up an org level Variable Set. Set org_id only.
data "harness_platform_infra_variable_set" "org_level" {
  identifier = "org_variable_set"
  org_id     = harness_platform_organization.example.id
}

# Look up a project level Variable Set. Set both org_id and project_id.
data "harness_platform_infra_variable_set" "project_level" {
  identifier = "project_variable_set"
  org_id     = harness_platform_organization.example.id
  project_id = harness_platform_project.example.id
}

# Consuming a Variable Set from a Workspace.
#
# The exported `id` is the bare identifier, with no scope prefix. Harness resolves an
# unprefixed reference against the Workspace's own org and project, so a Variable Set
# that lives above the Workspace must be referenced with a scope prefix:
#
#   account level -> "account.${...id}"
#   org level     -> "org.${...id}"
#   project level -> "${...id}" (no prefix, must be the same project as the Workspace)
#
# Referencing an account or org level Variable Set without the prefix fails with
# "404 Not Found ... variable set not found", because the lookup is scoped to the project.
resource "harness_platform_workspace" "example" {
  identifier          = "example"
  name                = "example"
  org_id              = harness_platform_organization.example.id
  project_id          = harness_platform_project.example.id
  provisioner_type    = "terraform"
  provisioner_version = "1.5.7"

  repository           = "https://github.com/org/repo"
  repository_branch    = "main"
  repository_path      = "tf/aws/basic"
  repository_connector = harness_platform_connector_github.example.id
  provider_connector   = harness_platform_connector_aws.example.id

  variable_sets = [
    "account.${data.harness_platform_infra_variable_set.account_level.id}",
    "org.${data.harness_platform_infra_variable_set.org_level.id}",
    data.harness_platform_infra_variable_set.project_level.id,
  ]
}
