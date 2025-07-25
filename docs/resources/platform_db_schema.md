---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_db_schema Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Harness DBDevOps Schema.
---

# harness_platform_db_schema (Resource)

Resource for creating a Harness DBDevOps Schema.

## Example Usage

```terraform
resource "harness_platform_db_schema" "test" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
  name       = "name"
  service    = "service1"
  tags       = ["foo:bar", "bar:foo"]
  schema_source {
    connector    = "gitConnector"
    repo         = "TestRepo"
    location     = "db/example-changelog.yaml"
    archive_path = "path/to/archive.zip"
  }
}

resource "harness_platform_db_schema" "test" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
  name       = "name"
  service    = "service1"
  type       = "Repository"
  tags       = ["foo:bar", "bar:foo"]
  schema_source {
    connector    = "gitConnector"
    repo         = "TestRepo"
    location     = "db/example-changelog.yaml"
    archive_path = "path/to/archive.zip"
  }
}


resource "harness_platform_db_schema" "test" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"
  name       = "name"
  service    = "service1"
  type       = "Script"
  tags       = ["foo:bar", "bar:foo"]
  changelog_script {
    image    = "plugins/image"
    command  = "echo \\\"hello dbops\\\""
    shell    = "sh/bash"
    location = "db/example-changelog.yaml"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.

### Optional

- `changelog_script` (Block List, Max: 1) Configuration to clone changeSets using script (see [below for nested schema](#nestedblock--changelog_script))
- `description` (String) Description of the resource.
- `schema_source` (Block List, Max: 1) Provides a connector and path at which to find the database schema representation (see [below for nested schema](#nestedblock--schema_source))
- `service` (String) The service associated with schema
- `tags` (Set of String) Tags to associate with the resource.
- `type` (String) Type of the database schema. Valid values are: SCRIPT, REPOSITORY

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--changelog_script"></a>
### Nested Schema for `changelog_script`

Optional:

- `command` (String) Script to clone changeSets
- `image` (String) The fully-qualified name (FQN) of the image
- `location` (String) Path to changeLog file
- `shell` (String) Type of the shell. For example Sh or Bash


<a id="nestedblock--schema_source"></a>
### Nested Schema for `schema_source`

Required:

- `connector` (String) Connector to repository at which to find details about the database schema
- `location` (String) The path within the specified repository at which to find details about the database schema

Optional:

- `archive_path` (String) If connector type is artifactory, path to the archive file which contains the changeLog
- `repo` (String) If connector url is of account, which repository to connect to using the connector

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Import project level db schema
terraform import harness_platform_db_schema.example <org_id>/<project_id>/<db_schema_id>
```
