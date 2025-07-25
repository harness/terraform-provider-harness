---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_template Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for retrieving a Harness pipeline.
---

# harness_platform_template (Data Source)

Data source for retrieving a Harness pipeline.

## Example Usage

```terraform
#For account level template
data "harness_platform_template" "example" {
  identifier = "identifier"
  version    = "version"
}

#For org level template
data "harness_platform_template" "example1" {
  identifier = "identifier"
  version    = "version"
  org_id     = "org_id"
}

#For project level template
data "harness_platform_template" "example2" {
  identifier = "identifier"
  version    = "version"
  org_id     = "org_id"
  project_id = "project_id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `branch_name` (String) Version Label for Template.
- `child_type` (String) Defines child template type.
- `git_details` (Block List, Max: 1) Contains parameters related to creating an Entity for Git Experience. (see [below for nested schema](#nestedblock--git_details))
- `identifier` (String) Unique identifier of the resource.
- `is_stable` (Boolean) True if given version for template to be set as stable.
- `name` (String) Name of the resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `scope` (String) Scope of template.
- `version` (String) Version Label for Template.

### Read-Only

- `connector_ref` (String) Identifier of the Harness Connector used for CRUD operations on the Entity.
- `description` (String) Description of the resource.
- `id` (String) The ID of this resource.
- `store_type` (String) Specifies whether the Entity is to be stored in Git or not. Possible values: INLINE, REMOTE.
- `tags` (Set of String) Tags to associate with the resource.
- `template_yaml` (String) Yaml for creating new Template.

<a id="nestedblock--git_details"></a>
### Nested Schema for `git_details`

Optional:

- `branch_name` (String) Name of the branch.
- `file_path` (String) File path of the Entity in the repository.
- `file_url` (String) File url of the Entity in the repository.
- `last_commit_id` (String) Last commit identifier (for Git Repositories other than Github). To be provided only when updating Pipeline.
- `last_object_id` (String) Last object identifier (for Github). To be provided only when updating Pipeline.
- `repo_name` (String) Name of the repository.
- `repo_url` (String) Repo url of the Entity in the repository.
