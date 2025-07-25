---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_docker Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Docker connector.
---

# harness_platform_connector_docker (Resource)

Resource for creating a Docker connector.

## Example Usage

```terraform
# credentials anonymous
resource "harness_platform_connector_docker" "test" {
  identifier  = "identifer"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  type               = "DockerHub"
  url                = "https://hub.docker.com"
  delegate_selectors = ["harness-delegate"]
}

# credentials username password
resource "harness_platform_connector_docker" "test" {
  identifier  = "identifer"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  type               = "DockerHub"
  url                = "https://hub.docker.com"
  delegate_selectors = ["harness-delegate"]
  credentials {
    username     = "admin"
    password_ref = "account.secret_id"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `type` (String) The type of the docker registry. Valid options are DockerHub, Harbor, Other, Quay
- `url` (String) The URL of the docker registry.

### Optional

- `credentials` (Block List, Max: 1) The credentials to use for the docker registry. If not specified then the connection is made to the registry anonymously. (see [below for nested schema](#nestedblock--credentials))
- `delegate_selectors` (Set of String) Tags to filter delegates for connection.
- `description` (String) Description of the resource.
- `execute_on_delegate` (Boolean) Execute on delegate or not.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `tags` (Set of String) Tags to associate with the resource.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--credentials"></a>
### Nested Schema for `credentials`

Required:

- `password_ref` (String) The reference to the Harness secret containing the password to use for the docker registry. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.

Optional:

- `username` (String) The username to use for the docker registry.
- `username_ref` (String) The reference to the Harness secret containing the username to use for the docker registry. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Import account level docker connector 
terraform import harness_platform_connector_docker.example <connector_id>

# Import org level docker connector 
terraform import harness_platform_connector_docker.example <ord_id>/<connector_id>

# Import project level docker connector 
terraform import harness_platform_connector_docker.example <org_id>/<project_id>/<connector_id>
```
