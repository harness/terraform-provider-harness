---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_pipeline Resource - terraform-provider-harness"
subcategory: ""
description: |-
  Resource for creating a Harness pipeline.
---

# harness_platform_pipeline (Resource)

Resource for creating a Harness pipeline.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `yaml` (String) YAML of the pipeline.

### Optional

- `description` (String) Description of the resource.
- `id` (String) The ID of this resource.
- `tags` (Set of String) Tags to associate with the resource. Tags should be in the form `name:value`.


