---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_azure_artifacts Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating an Azure Artifacts connector.
---

# harness_platform_connector_azure_artifacts (Resource)

Resource for creating an Azure Artifacts connector.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `credentials` (Block List, Min: 1, Max: 1) Credentials to use for authentication. (see [below for nested schema](#nestedblock--credentials))
- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `url` (String) URL of the Azure Artifacts server.

### Optional

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

- `token_ref` (String) Reference to a secret containing the token to use for authentication. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}.
