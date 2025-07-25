---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_sumologic Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Sumologic connector.
---

# harness_platform_connector_sumologic (Resource)

Resource for creating a Sumologic connector.

## Example Usage

```terraform
resource "harness_platform_connector_sumologic" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://api.us2.sumologic.com/"
  delegate_selectors = ["harness-delegate"]
  access_id_ref      = "account.secret_id"
  access_key_ref     = "account.secret_id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `access_id_ref` (String) Reference to the Harness secret containing the access id. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.
- `access_key_ref` (String) Reference to the Harness secret containing the access key. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.
- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `url` (String) URL of the SumoLogic server.

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
# Import account level sumologic connector 
terraform import harness_platform_connector_sumologic.example <connector_id>

# Import org level sumologic connector 
terraform import harness_platform_connector_sumologic.example <ord_id>/<connector_id>

# Import project level sumologic connector 
terraform import harness_platform_connector_sumologic.example <org_id>/<project_id>/<connector_id>
```
