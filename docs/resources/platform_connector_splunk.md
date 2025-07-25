---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_splunk Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Splunk connector.
---

# harness_platform_connector_splunk (Resource)

Resource for creating a Splunk connector.

## Example Usage

```terraform
resource "harness_platform_connector_splunk" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://splunk.com/"
  delegate_selectors = ["harness-delegate"]
  account_id         = "splunk_account_id"
  username           = "username"
  password_ref       = "account.secret_id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Splunk account id.
- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `password_ref` (String) The reference to the Harness secret containing the Splunk password. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.
- `url` (String) URL of the Splunk server.
- `username` (String) The username used for connecting to Splunk.

### Optional

- `delegate_selectors` (Set of String) Tags to filter delegates for connection.
- `description` (String) Description of the resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `tags` (Set of String) Tags to associate with the resource.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Import account level splunk connector 
terraform import harness_platform_connector_splunk.example <connector_id>

# Import org level splunk connector 
terraform import harness_platform_connector_splunk.example <ord_id>/<connector_id>

# Import project level splunk connector 
terraform import harness_platform_connector_splunk.example <org_id>/<project_id>/<connector_id>
```
