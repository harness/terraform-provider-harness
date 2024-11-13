---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_infra_module Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for retrieving modules from the module registry.
---

# harness_platform_infra_module (Data Source)

Data source for retrieving modules from the module registry.

## Example Usage

```terraform
data "harness_platform_infra_module" "test" {
  identifier = "identifier"
  name       = "name"
  system     = "system"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Identifier of the module
- `name` (String) Name of the module
- `system` (String) Provider of the module

### Optional

- `created` (Number) Timestamp when the module was created
- `description` (String) Description of the module
- `repository` (String) For account connectors, the repository where the module is stored
- `repository_branch` (String) Repository Branch in which the module should be accessed
- `repository_commit` (String) Repository Commit in which the module should be accessed
- `repository_connector` (String) Repository Connector is the reference to the connector for the repository
- `repository_path` (String) Repository Path is the path in which the module resides
- `repository_url` (String) URL where the module is stored
- `synced` (Number) Timestamp when the module was last synced
- `tags` (String) Tags associated with the module