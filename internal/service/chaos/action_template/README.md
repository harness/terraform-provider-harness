# Chaos Action Template Resource

## Overview

The `harness_chaos_action_template` resource manages Chaos Action Templates in Harness. Action templates define reusable actions that can be used in chaos experiments to perform various operations during fault injection.

## Features

- ✅ Create, read, update, and delete action templates
- ✅ Support for account, org, and project scopes
- ✅ Association with chaos hubs
- ✅ Import existing action templates
- ✅ Data source for reading action templates by identity or name
- ✅ Enhanced error handling with detailed API error messages

## Resource: `harness_chaos_action_template`

### Example Usage

#### Project-Level Action Template

```hcl
resource "harness_chaos_action_template" "example" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "custom_hub"
  
  identity    = "my_action_template"
  name        = "My Action Template"
  description = "Custom action template for chaos experiments"
  
  tags = ["production", "critical"]
  type = "HTTP"
}
```

#### Org-Level Action Template

```hcl
resource "harness_chaos_action_template" "org_level" {
  org_id       = "my_org"
  hub_identity = "org_hub"
  
  identity    = "org_action_template"
  name        = "Org Action Template"
  description = "Organization-wide action template"
}
```

#### Account-Level Action Template

```hcl
resource "harness_chaos_action_template" "account_level" {
  hub_identity = "account_hub"
  
  identity    = "account_action_template"
  name        = "Account Action Template"
  description = "Account-wide action template"
}
```

### Argument Reference

#### Required Arguments

- `identity` - (String, ForceNew) Unique identifier for the action template. This is immutable after creation.
- `name` - (String) Name of the action template.
- `hub_identity` - (String, ForceNew) Identity of the chaos hub this action template belongs to.

#### Optional Arguments

- `org_id` - (String, ForceNew) Organization identifier. Required for org and project-level templates.
- `project_id` - (String, ForceNew) Project identifier. Required for project-level templates.
- `description` - (String) Description of the action template.
- `tags` - (List of String) Tags to associate with the action template.
- `type` - (String) Type of the action template (e.g., HTTP, Kubernetes, Shell).
- `infrastructure_type` - (String) Infrastructure type for the action template (e.g., Kubernetes, Linux, Windows).

### Attribute Reference

In addition to the arguments above, the following attributes are exported:

- `id` - The ID of the action template in the format: `account_id/org_id/project_id/hub_identity/identity`
- `account_id` - Account identifier.
- `id_internal` - Internal ID of the action template.
- `revision` - Revision number of the action template.
- `is_default` - Whether this is the default version for predefined actions.
- `is_enterprise` - Whether this is an enterprise action template.
- `is_removed` - Whether the action template has been removed.
- `created_at` - Creation timestamp (Unix epoch).
- `updated_at` - Last update timestamp (Unix epoch).
- `created_by` - User who created the action template.
- `updated_by` - User who last updated the action template.
- `template` - Template content/definition.

### Import

Action templates can be imported using one of the following formats:

```bash
# Project level
terraform import harness_chaos_action_template.example org_id/project_id/hub_identity/identity

# Org level
terraform import harness_chaos_action_template.example org_id/hub_identity/identity

# Account level
terraform import harness_chaos_action_template.example hub_identity/identity
```

Example:

```bash
terraform import harness_chaos_action_template.example my_org/my_project/custom_hub/my_action_template
```

## Data Source: `harness_chaos_action_template`

### Example Usage

#### Lookup by Identity

```hcl
data "harness_chaos_action_template" "example" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "custom_hub"
  identity     = "my_action_template"
}

output "action_template_revision" {
  value = data.harness_chaos_action_template.example.revision
}
```

#### Lookup by Name

```hcl
data "harness_chaos_action_template" "example" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "custom_hub"
  name         = "My Action Template"
}

output "action_template_identity" {
  value = data.harness_chaos_action_template.example.identity
}
```

### Argument Reference

#### Required Arguments

- `hub_identity` - (String) Identity of the chaos hub to search in.

#### Optional Arguments

- `org_id` - (String) Organization identifier.
- `project_id` - (String) Project identifier.
- `identity` - (String) Unique identifier of the action template. Either `identity` or `name` must be specified.
- `name` - (String) Name of the action template. Either `identity` or `name` must be specified.

**Note:** You must specify either `identity` or `name`, but not both.

### Attribute Reference

All attributes from the resource are available in the data source.

## Scope Hierarchy

Action templates support three levels of scope:

1. **Account Level**: No `org_id` or `project_id` specified
   - Accessible across the entire account
   - Managed by account administrators

2. **Org Level**: `org_id` specified, no `project_id`
   - Accessible within the organization
   - Managed by organization administrators

3. **Project Level**: Both `org_id` and `project_id` specified
   - Accessible only within the project
   - Managed by project members

## Hub Association

All action templates must be associated with a chaos hub via the `hub_identity` field. The hub must exist before creating the action template.

Example with hub creation:

```hcl
resource "harness_chaos_hub_v2" "custom" {
  org_id      = "my_org"
  project_id  = "my_project"
  identity    = "custom_hub"
  name        = "Custom Chaos Hub"
  description = "Custom hub for action templates"
}

resource "harness_chaos_action_template" "example" {
  org_id       = harness_chaos_hub_v2.custom.org_id
  project_id   = harness_chaos_hub_v2.custom.project_id
  hub_identity = harness_chaos_hub_v2.custom.identity
  
  identity = "my_action"
  name     = "My Action"
}
```

## Error Handling

This resource uses enhanced error handling that provides detailed API error messages:

- **Create/Update/Delete**: Uses `HandleChaosApiError()` for detailed error messages
- **Read**: Uses `HandleChaosReadApiError()` which gracefully handles 404 errors
- **Import**: Validates import ID format and provides helpful error messages

Example error output:

```
Error: Internal Server Error: error occurred while creating action template: "action template already exists", statuscode 500
```

## Best Practices

1. **Use Descriptive Identities**: Choose meaningful, immutable identities for your action templates
2. **Leverage Tags**: Use tags to categorize and organize action templates
3. **Version Control**: Use revision numbers to track changes
4. **Hub Organization**: Group related action templates in the same hub
5. **Scope Appropriately**: Use the most restrictive scope that meets your needs

## Limitations

- `identity` and `hub_identity` are immutable (ForceNew) - changing them will destroy and recreate the resource
- Scope fields (`org_id`, `project_id`) are immutable (ForceNew)
- Action templates must be associated with an existing chaos hub

## Related Resources

- `harness_chaos_hub_v2` - Manage chaos hubs
- `harness_chaos_probe_template` - Manage probe templates
- `harness_chaos_fault_template` - Manage fault templates
- `harness_chaos_experiment_template` - Manage experiment templates

## API Reference

This resource uses the Harness Chaos REST API:

- **Create**: `POST /rest/hubs/{hubIdentity}/action-templates`
- **Read**: `GET /rest/hubs/{hubIdentity}/action-templates/{identity}`
- **Update**: `PUT /rest/hubs/{hubIdentity}/action-templates/{identity}`
- **Delete**: `DELETE /rest/hubs/{hubIdentity}/action-templates/{identity}`
- **List**: `GET /rest/hubs/{hubIdentity}/action-templates`

## Support

For issues or questions:
- GitHub Issues: [terraform-provider-harness](https://github.com/harness/terraform-provider-harness/issues)
- Harness Documentation: [Chaos Engineering](https://docs.harness.io/category/chaos-engineering)
