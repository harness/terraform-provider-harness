---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_aws_secret_manager Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating an AWS Secret Manager connector.
---

# harness_platform_connector_aws_secret_manager (Resource)

Resource for creating an AWS Secret Manager connector.

## Example Usage

```terraform
# Credentials inherit_from_delegate
resource "harness_platform_connector_aws_secret_manager" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  secret_name_prefix = "test"
  region             = "us-east-1"
  delegate_selectors = ["harness-delegate"]
  use_put_secret     = false
  credentials {
    inherit_from_delegate = true
  }
}

# Credentials manual
resource "harness_platform_connector_aws_secret_manager" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  secret_name_prefix = "test"
  region             = "us-east-1"
  delegate_selectors = ["harness-delegate"]
  use_put_secret     = false
  credentials {
    manual {
      secret_key_ref = "account.secret_id"
      access_key_ref = "account.secret_id"
    }
  }
}

# Credentials assume_role
resource "harness_platform_connector_aws_secret_manager" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  secret_name_prefix = "test"
  region             = "us-east-1"
  delegate_selectors = ["harness-delegate"]
  default            = true
  use_put_secret     = false
  credentials {
    assume_role {
      role_arn    = "somerolearn"
      external_id = "externalid"
      duration    = 900
    }
  }
}

# Credentials oidc using Harness Platform
resource "harness_platform_connector_aws_secret_manager" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  secret_name_prefix  = "test"
  region              = "us-east-1"
  default             = true
  use_put_secret      = false
  execute_on_delegate = false

  credentials {
    oidc_authentication {
      iam_role_arn = "arn:aws:iam:testarn"
    }
  }
}

# Force delete true
resource "harness_platform_connector_aws_secret_manager" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  secret_name_prefix            = "test"
  region                        = "us-east-1"
  delegate_selectors            = ["harness-delegate"]
  default                       = true
  force_delete_without_recovery = true
  credentials {
    assume_role {
      role_arn    = "somerolearn"
      external_id = "externalid"
      duration    = 900
    }
  }
}

# Credentials oidc using Delegate
resource "harness_platform_connector_aws_secret_manager" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  secret_name_prefix = "test"
  region             = "us-east-1"
  delegate_selectors = ["harness-delegate"]
  default            = true
  use_put_secret     = false

  credentials {
    oidc_authentication {
      iam_role_arn = "arn:aws:iam:testarn"
    }
  }
}

# With recovery duration of 15 days
resource "harness_platform_connector_aws_secret_manager" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  secret_name_prefix      = "test"
  region                  = "us-east-1"
  delegate_selectors      = ["harness-delegate"]
  default                 = true
  recovery_window_in_days = 15
  credentials {
    assume_role {
      role_arn    = "somerolearn"
      external_id = "externalid"
      duration    = 900
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `credentials` (Block List, Min: 1, Max: 1) Credentials to connect to AWS. (see [below for nested schema](#nestedblock--credentials))
- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `region` (String) The AWS region where the AWS Secret Manager is.

### Optional

- `default` (Boolean) Use as Default Secrets Manager.
- `delegate_selectors` (Set of String) Tags to filter delegates for connection.
- `description` (String) Description of the resource.
- `execute_on_delegate` (Boolean) Run the operation on the delegate or harness platform.
- `force_delete_without_recovery` (Boolean) Whether to force delete secret value or not.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `recovery_window_in_days` (Number) Recovery duration in days in AWS Secrets Manager.
- `secret_name_prefix` (String) A prefix to be added to all secrets.
- `tags` (Set of String) Tags to associate with the resource.
- `use_put_secret` (Boolean) Whether to update secret value using putSecretValue action.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--credentials"></a>
### Nested Schema for `credentials`

Optional:

- `assume_role` (Block List, Max: 1) Connect using STS assume role. (see [below for nested schema](#nestedblock--credentials--assume_role))
- `inherit_from_delegate` (Boolean) Inherit the credentials from from the delegate.
- `manual` (Block List, Max: 1) Specify the AWS key and secret used for authenticating. (see [below for nested schema](#nestedblock--credentials--manual))
- `oidc_authentication` (Block List, Max: 1) Authentication using harness oidc. (see [below for nested schema](#nestedblock--credentials--oidc_authentication))

<a id="nestedblock--credentials--assume_role"></a>
### Nested Schema for `credentials.assume_role`

Required:

- `duration` (Number) The duration, in seconds, of the role session. The value can range from 900 seconds (15 minutes) to 3600 seconds (1 hour). By default, the value is set to 3600 seconds. An expiration can also be specified in the client request body. The minimum value is 1 hour.
- `role_arn` (String) The ARN of the role to assume.

Optional:

- `external_id` (String) If the administrator of the account to which the role belongs provided you with an external ID, then enter that value.


<a id="nestedblock--credentials--manual"></a>
### Nested Schema for `credentials.manual`

Required:

- `secret_key_ref` (String) The reference to the Harness secret containing the AWS secret key. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.

Optional:

- `access_key_plain_text` (String) The plain text AWS access key.
- `access_key_ref` (String) The reference to the Harness secret containing the AWS access key. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.


<a id="nestedblock--credentials--oidc_authentication"></a>
### Nested Schema for `credentials.oidc_authentication`

Required:

- `iam_role_arn` (String) The IAM role ARN.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Import account level aws secret manager connector 
terraform import harness_platform_connector_aws_secret_manager.example <connector_id>

# Import org level aws secret manager connector 
terraform import harness_platform_connector_aws_secret_manager.example <ord_id>/<connector_id>

# Import project level aws secret manager connector 
terraform import harness_platform_connector_aws_secret_manager.example <org_id>/<project_id>/<connector_id>
```
