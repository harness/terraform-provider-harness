---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_db_schema Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for retrieving a Harness DBDevOps Schema.
---

# harness_platform_db_schema (Data Source)

Data source for retrieving a Harness DBDevOps Schema.

## Example Usage

```terraform
data "harness_platform_db_schema" "example" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.

### Optional

- `name` (String) Name of the resource.

### Read-Only

- `description` (String) Description of the resource.
- `id` (String) The ID of this resource.
- `schema_source` (List of Object) Provides a connector and path at which to find the database schema representation (see [below for nested schema](#nestedatt--schema_source))
- `service` (String) The service associated with schema
- `tags` (Set of String) Tags to associate with the resource.

<a id="nestedatt--schema_source"></a>
### Nested Schema for `schema_source`

Read-Only:

- `connector` (String)
- `location` (String)
- `repo` (String)