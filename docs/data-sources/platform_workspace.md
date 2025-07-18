---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_workspace Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for retrieving workspaces.
---

# harness_platform_workspace (Data Source)

Data source for retrieving workspaces.

## Example Usage

```terraform
data "harness_platform_workspace" "test" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Identifier of the Workspace
- `org_id` (String) Organization Identifier
- `project_id` (String) Project Identifier

### Optional

- `connector` (Block Set) Provider connectors configured on the Workspace. Only one connector of a type is supported (see [below for nested schema](#nestedblock--connector))
- `description` (String) Description of the Workspace
- `repository_branch` (String) Repository Branch in which the code should be accessed
- `repository_commit` (String) Repository Tag in which the code should be accessed
- `repository_sha` (String) Repository Commit SHA in which the code should be accessed
- `tags` (Set of String) Tags to associate with the resource.
- `variable_sets` (List of String) Variable sets to use.

### Read-Only

- `cost_estimation_enabled` (Boolean) If enabled cost estimation operations will be performed in this workspace
- `default_pipelines` (Map of String) Default pipelines associated with this workspace
- `environment_variable` (Block Set) Environment variables configured on the workspace (see [below for nested schema](#nestedblock--environment_variable))
- `id` (String) The ID of this resource.
- `name` (String) Name of the Workspace
- `provider_connector` (String) Provider Connector is the reference to the connector for the infrastructure provider
- `provisioner_type` (String) Provisioner type defines the provisioning tool to use.
- `provisioner_version` (String) Provisioner Version defines the tool version to use
- `repository` (String) Repository is the name of the repository to use
- `repository_connector` (String) Repository Connector is the reference to the connector to use for this code
- `repository_path` (String) Repository Path is the path in which the infra code resides
- `terraform_variable` (Block Set) Terraform variables configured on the workspace (see [below for nested schema](#nestedblock--terraform_variable))
- `terraform_variable_file` (Block Set) Terraform variables files configured on the workspace (see [below for nested schema](#nestedblock--terraform_variable_file))

<a id="nestedblock--connector"></a>
### Nested Schema for `connector`

Required:

- `connector_ref` (String) Connector Ref is the reference to the connector
- `type` (String) Type is the connector type of the connector. Supported types: aws, azure, gcp


<a id="nestedblock--environment_variable"></a>
### Nested Schema for `environment_variable`

Read-Only:

- `key` (String) Key is the identifier for the variable`
- `value` (String) value is the value of the variable
- `value_type` (String) Value type indicates the value type of the variable, text or secret


<a id="nestedblock--terraform_variable"></a>
### Nested Schema for `terraform_variable`

Read-Only:

- `key` (String) Key is the identifier for the variable`
- `value` (String) value is the value of the variable
- `value_type` (String) Value type indicates the value type of the variable, text or secret


<a id="nestedblock--terraform_variable_file"></a>
### Nested Schema for `terraform_variable_file`

Read-Only:

- `repository` (String) Repository is the name of the repository to fetch the code from.
- `repository_branch` (String) Repository branch is the name of the branch to fetch the variables from. This cannot be set if repository commit or sha is set
- `repository_commit` (String) Repository commit is tag to fetch the variables from. This cannot be set if repository branch or sha is set.
- `repository_connector` (String) Repository connector is the reference to the connector used to fetch the variables.
- `repository_path` (String) Repository path is the path in which the variables reside.
- `repository_sha` (String) Repository commit is SHA to fetch the variables from. This cannot be set if repository branch or commit is set.
