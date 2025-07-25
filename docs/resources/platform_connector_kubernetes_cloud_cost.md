---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_kubernetes_cloud_cost Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Kubernetes Cloud Cost connector.
---

# harness_platform_connector_kubernetes_cloud_cost (Resource)

Resource for creating a Kubernetes Cloud Cost connector.

## Example Usage

```terraform
resource "harness_platform_connector_kubernetes_cloud_cost" "example" {
  identifier  = "identifier"
  name        = "name"
  description = "example"
  tags        = ["foo:bar"]

  features_enabled = ["VISIBILITY", "OPTIMIZATION"]
  connector_ref    = "connector_ref"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connector_ref` (String) Reference of the Connector. To reference a connector at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a connector at the account scope, prefix 'account` to the expression: account.{identifier}.
- `features_enabled` (Set of String) Indicates which feature to enable among Billing, Optimization, and Visibility.
- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.

### Optional

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
# Import account level kubernetes cloud cost connector 
terraform import harness_platform_connector_kubernetes_cloud_cost.example <connector_id>

# Import org level kubernetes cloud cost connector 
terraform import harness_platform_connector_kubernetes_cloud_cost.example <ord_id>/<connector_id>

# Import project level kubernetes cloud cost connector 
terraform import harness_platform_connector_kubernetes_cloud_cost.example <org_id>/<project_id>/<connector_id>
```
