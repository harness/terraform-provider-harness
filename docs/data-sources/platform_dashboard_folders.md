---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_dashboard_folders Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for retrieving a Harness Dashboard Folder.
---

# harness_platform_dashboard_folders (Data Source)

Data source for retrieving a Harness Dashboard Folder.

## Example Usage

```terraform
data "harness_platform_dashboard_folders" "folder" {
  id = "id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Identifier of the folder.

### Optional

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.

### Read-Only

- `created_at` (String) Created DateTime of the folder.
- `description` (String) Description of the resource.
- `tags` (Set of String) Tags to associate with the resource.
