---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_ccm_filters Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for retrieving a Harness CCM Filter.
---

# harness_platform_ccm_filters (Data Source)

Data source for retrieving a Harness CCM Filter.

## Example Usage

```terraform
data "harness_platform_ccm_filters" "test" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
  type       = "CCMRecommendation"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `type` (String) Type of filter.

### Optional

- `org_id` (String) Organization Identifier for the Entity.
- `project_id` (String) Project Identifier for the Entity.

### Read-Only

- `filter_properties` (List of Object) Properties of the filter entity defined in Harness. (see [below for nested schema](#nestedatt--filter_properties))
- `filter_visibility` (String) This indicates visibility of filter. By default, everyone can view this filter.
- `id` (String) The ID of this resource.
- `name` (String) Name of the Filter.

<a id="nestedatt--filter_properties"></a>
### Nested Schema for `filter_properties`

Read-Only:

- `filter_type` (String)
- `tags` (Set of String)
