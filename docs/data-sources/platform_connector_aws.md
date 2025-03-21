---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_aws Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Datasource for looking up an AWS connector.
---

# harness_platform_connector_aws (Data Source)

Datasource for looking up an AWS connector.

## Example Usage

```terraform
data "harness_platform_connector_aws" "example" {
  identifier = "identifier"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.

### Optional

- `name` (String) Name of the resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.

### Read-Only

- `cross_account_access` (List of Object) Select this option if you want to use one AWS account for the connection, but you want to deploy or build in a different AWS account. In this scenario, the AWS account used for AWS access in Credentials will assume the IAM role you specify in Cross-account role ARN setting. This option uses the AWS Security Token Service (STS) feature. (see [below for nested schema](#nestedatt--cross_account_access))
- `description` (String) Description of the resource.
- `equal_jitter_backoff_strategy` (List of Object) Equal Jitter BackOff Strategy. (see [below for nested schema](#nestedatt--equal_jitter_backoff_strategy))
- `execute_on_delegate` (Boolean) Execute on delegate or not.
- `fixed_delay_backoff_strategy` (List of Object) Fixed Delay BackOff Strategy. (see [below for nested schema](#nestedatt--fixed_delay_backoff_strategy))
- `full_jitter_backoff_strategy` (List of Object) Full Jitter BackOff Strategy. (see [below for nested schema](#nestedatt--full_jitter_backoff_strategy))
- `id` (String) The ID of this resource.
- `inherit_from_delegate` (List of Object) Inherit credentials from the delegate. (see [below for nested schema](#nestedatt--inherit_from_delegate))
- `irsa` (List of Object) Use IAM role for service accounts. (see [below for nested schema](#nestedatt--irsa))
- `manual` (List of Object) Use IAM role for service accounts. (see [below for nested schema](#nestedatt--manual))
- `oidc_authentication` (List of Object) Authentication using harness oidc. (see [below for nested schema](#nestedatt--oidc_authentication))
- `tags` (Set of String) Tags to associate with the resource.

<a id="nestedatt--cross_account_access"></a>
### Nested Schema for `cross_account_access`

Read-Only:

- `external_id` (String)
- `role_arn` (String)


<a id="nestedatt--equal_jitter_backoff_strategy"></a>
### Nested Schema for `equal_jitter_backoff_strategy`

Read-Only:

- `base_delay` (Number)
- `max_backoff_time` (Number)
- `retry_count` (Number)


<a id="nestedatt--fixed_delay_backoff_strategy"></a>
### Nested Schema for `fixed_delay_backoff_strategy`

Read-Only:

- `fixed_backoff` (Number)
- `retry_count` (Number)


<a id="nestedatt--full_jitter_backoff_strategy"></a>
### Nested Schema for `full_jitter_backoff_strategy`

Read-Only:

- `base_delay` (Number)
- `max_backoff_time` (Number)
- `retry_count` (Number)


<a id="nestedatt--inherit_from_delegate"></a>
### Nested Schema for `inherit_from_delegate`

Read-Only:

- `delegate_selectors` (Set of String)
- `region` (String)


<a id="nestedatt--irsa"></a>
### Nested Schema for `irsa`

Read-Only:

- `delegate_selectors` (Set of String)
- `region` (String)


<a id="nestedatt--manual"></a>
### Nested Schema for `manual`

Read-Only:

- `access_key` (String)
- `access_key_plain_text` (String)
- `access_key_ref` (String)
- `delegate_selectors` (Set of String)
- `region` (String)
- `secret_key_ref` (String)
- `session_token_ref` (String)


<a id="nestedatt--oidc_authentication"></a>
### Nested Schema for `oidc_authentication`

Read-Only:

- `delegate_selectors` (Set of String)
- `iam_role_arn` (String)
- `region` (String)
