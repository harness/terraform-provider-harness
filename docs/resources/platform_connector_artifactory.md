---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_artifactory Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating an Artifactory connector.
---

# harness_platform_connector_artifactory (Resource)

Resource for creating an Artifactory connector.

### References:
- For details on how to onboard with Terraform, please see [Harness Terraform Provider Overview](https://developer.harness.io/docs/platform/automation/terraform/harness-terraform-provider-overview/)
- To understand how to use the Connectors, please see [Documentation](https://developer.harness.io/docs/category/connectors)

## Example to create Artifactory Connector at different levels (Org, Project, Account)
### Account Level
```terraform
# Authentication mechanism as username and password
resource "harness_platform_connector_artifactory" "example" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  
  url                = "https://artifactory.example.com"
  delegate_selectors = ["harness-delegate"]
  credentials {
    username     = "admin"
    password_ref = "account.secret_id"
  }
}

# Authentication mechanism as anonymous
resource "harness_platform_connector_artifactory" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://artifactory.example.com"
  delegate_selectors = ["harness-delegate"]
}
```

### Org Level
```terraform
# Authentication mechanism as username and password
resource "harness_platform_connector_artifactory" "example" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  org_id      = harness_platform_project.test.org_id
  
  url                = "https://artifactory.example.com"
  delegate_selectors = ["harness-delegate"]
  credentials {
    username     = "admin"
    password_ref = "account.secret_id"
  }
}

# Authentication mechanism as anonymous
resource "harness_platform_connector_artifactory" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  org_id      = harness_platform_project.test.org_id

  url                = "https://artifactory.example.com"
  delegate_selectors = ["harness-delegate"]
}
```

### Project Level
```terraform
# Authentication mechanism as username and password
resource "harness_platform_connector_artifactory" "example" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  org_id      = harness_platform_project.test.org_id
  project_id  = harness_platform_project.test.id

  url                = "https://artifactory.example.com"
  delegate_selectors = ["harness-delegate"]
  credentials {
    username     = "admin"
    password_ref = "account.secret_id"
  }
}

# Authentication mechanism as anonymous
resource "harness_platform_connector_artifactory" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  org_id      = harness_platform_project.test.org_id
  project_id  = harness_platform_project.test.id

  url                = "https://artifactory.example.com"
  delegate_selectors = ["harness-delegate"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `url` (String) URL of the Artifactory server.

### Optional

- `credentials` (Block List, Max: 1) Credentials to use for authentication. (see [below for nested schema](#nestedblock--credentials))
- `delegate_selectors` (Set of String) Tags to filter delegates for connection.
- `description` (String) Description of the resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `tags` (Set of String) Tags to associate with the resource.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--credentials"></a>
### Nested Schema for `credentials`

Required:

- `password_ref` (String) Reference to a secret containing the password to use for authentication. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.

Optional:

- `username` (String) Username to use for authentication.
- `username_ref` (String) Reference to a secret containing the username to use for authentication. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Import account level artifactory connector 
terraform import harness_platform_connector_artifactory.example <connector_id>

# Import org level artifactory connector 
terraform import harness_platform_connector_artifactory.example <ord_id>/<connector_id>

# Import project level artifactory connector 
terraform import harness_platform_connector_artifactory.example <org_id>/<project_id>/<connector_id>
```
