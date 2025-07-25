---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_ssh_credential Data Source - terraform-provider-harness"
subcategory: "First Gen"
description: |-
  Data source for retrieving an SSH credential.
---

# harness_ssh_credential (Data Source)

Data source for retrieving an SSH credential.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) Unique identifier of the secret manager
- `name` (String) The name of the secret manager
- `usage_scope` (Block Set) This block is used for scoping the resource to a specific set of applications or environments. (see [below for nested schema](#nestedblock--usage_scope))

<a id="nestedblock--usage_scope"></a>
### Nested Schema for `usage_scope`

Optional:

- `application_id` (String) Id of the application to scope to. If empty then this scope applies to all applications.
- `environment_filter_type` (String) Type of environment filter applied. Cannot be used with `environment_id`. Valid options are NON_PRODUCTION_ENVIRONMENTS, PRODUCTION_ENVIRONMENTS.
- `environment_id` (String) Id of the id of the specific environment to scope to. Cannot be used with `environment_filter_type`.
