---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_gitlab Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Gitlab connector.
---

# harness_platform_connector_gitlab (Resource)

Resource for creating a Gitlab connector.

## Example Usage

```terraform
# Credentials http
resource "harness_platform_connector_gitlab" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://gitlab.com/account"
  connection_type    = "Account"
  validation_repo    = "some_repo"
  delegate_selectors = ["harness-delegate"]
  credentials {
    http {
      username  = "username"
      token_ref = "account.secret_id"
    }
  }
}

resource "harness_platform_connector_gitlab" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://gitlab.com/account"
  connection_type    = "Account"
  validation_repo    = "some_repo"
  delegate_selectors = ["harness-delegate"]
  credentials {
    http {
      username  = "username"
      token_ref = "account.secret_id"
    }
  }
  api_authentication {
    token_ref = "account.secret_id"
  }
}

# Credentials ssh
resource "harness_platform_connector_gitlab" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://gitlab.com/account"
  connection_type    = "Account"
  validation_repo    = "some_repo"
  delegate_selectors = ["harness-delegate"]
  credentials {
    ssh {
      ssh_key_ref = "account.test"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connection_type` (String) Whether the connection we're making is to a gitlab repository or a gitlab account. Valid values are Account, Repo.
- `credentials` (Block List, Min: 1, Max: 1) Credentials to use for the connection. (see [below for nested schema](#nestedblock--credentials))
- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `url` (String) URL of the gitlab repository or account.

### Optional

- `api_authentication` (Block List, Max: 1) Configuration for using the gitlab api. API Access is required for using “Git Experience”, for creation of Git based triggers, Webhooks management and updating Git statuses. (see [below for nested schema](#nestedblock--api_authentication))
- `delegate_selectors` (Set of String) Tags to filter delegates for connection.
- `description` (String) Description of the resource.
- `execute_on_delegate` (Boolean) Execute on delegate or not.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `tags` (Set of String) Tags to associate with the resource.
- `validation_repo` (String) Repository to test the connection with. This is only used when `connection_type` is `Account`.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--credentials"></a>
### Nested Schema for `credentials`

Optional:

- `http` (Block List, Max: 1) Authenticate using Username and password over http(s) for the connection. (see [below for nested schema](#nestedblock--credentials--http))
- `ssh` (Block List, Max: 1) Authenticate using SSH for the connection. (see [below for nested schema](#nestedblock--credentials--ssh))

<a id="nestedblock--credentials--http"></a>
### Nested Schema for `credentials.http`

Optional:

- `password_ref` (String) Reference to a secret containing the password to use for authentication. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.
- `token_ref` (String) Reference to a secret containing the personal access to use for authentication. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.
- `username` (String) Username to use for authentication.
- `username_ref` (String) Reference to a secret containing the username to use for authentication. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.


<a id="nestedblock--credentials--ssh"></a>
### Nested Schema for `credentials.ssh`

Required:

- `ssh_key_ref` (String) Reference to the Harness secret containing the ssh key. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.



<a id="nestedblock--api_authentication"></a>
### Nested Schema for `api_authentication`

Required:

- `token_ref` (String) Personal access token for interacting with the gitlab api. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Import account level gitlab connector 
terraform import harness_platform_connector_gitlab.example <connector_id>

# Import org level gitlab connector 
terraform import harness_platform_connector_gitlab.example <ord_id>/<connector_id>

# Import project level gitlab connector 
terraform import harness_platform_connector_gitlab.example <org_id>/<project_id>/<connector_id>
```
