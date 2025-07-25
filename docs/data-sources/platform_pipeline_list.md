---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_pipeline_list Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for retieving the Harness pipleine List
---

# harness_platform_pipeline_list (Data Source)

Data source for retieving the Harness pipleine List



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.

### Optional

- `identifier` (String) Unique identifier of the resource.
- `limit` (Number)
- `name` (String) Name of the resource.
- `page` (Number)

### Read-Only

- `description` (String) Description of the resource.
- `id` (String) The ID of this resource.
- `pipelines` (List of Object) (see [below for nested schema](#nestedatt--pipelines))
- `tags` (Set of String) Tags to associate with the resource.

<a id="nestedatt--pipelines"></a>
### Nested Schema for `pipelines`

Read-Only:

- `identifier` (String)
- `name` (String)
