---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_permissions Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for retrieving permissions.
---

# harness_platform_permissions (Data Source)

Data source for retrieving permissions.

## Example Usage

```terraform
data "harness_platform_permissions" "test" {
  org_id     = "org_id"
  project_id = "project_id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `org_id` (String) Organization Identifier
- `project_id` (String) Project Identifier

### Read-Only

- `id` (String) The ID of this resource.
- `permissions` (Set of Object) Response of the api (see [below for nested schema](#nestedatt--permissions))

<a id="nestedatt--permissions"></a>
### Nested Schema for `permissions`

Read-Only:

- `action` (String)
- `allowed_scope_levels` (Set of String)
- `identifier` (String)
- `include_in_all_roles` (Boolean)
- `name` (String)
- `resource_type` (String)
- `status` (String)
