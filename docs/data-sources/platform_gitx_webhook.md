---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_gitx_webhook Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Harness pipeline.
---

# harness_platform_gitx_webhook (Data Source)

Resource for creating a Harness pipeline.

## Example Usage

```terraform
//Create Project level webhook
resource "harness_platform_gitx_webhook" "test" {
  identifier = "webhook Identifier"
  name       = "webhook name"
  project_id = "projectIdentifier"
  org_id     = "orgIdentifier"
}

//Create Org level webhook
resource "harness_platform_gitx_webhook" "test" {
  identifier = "webhook Identifier"
  name       = "webhook name"
  org_id     = "orgIdentifier"
}

//Create Account level webhook
resource "harness_platform_gitx_webhook" "test" {
  identifier = "webhook Identifier"
  name       = "webhook name"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.

### Optional

- `description` (String) Description of the resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `tags` (Set of String) Tags to associate with the resource.

### Read-Only

- `id` (String) The ID of this resource.
