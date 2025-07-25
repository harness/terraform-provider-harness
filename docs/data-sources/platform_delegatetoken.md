---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_delegatetoken Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for retrieving a Harness delegate Token.
---

# harness_platform_delegatetoken (Data Source)

Data source for retrieving a Harness delegate Token.

## Example Usage

```terraform
# Look up a delegate token at account level by name
data "harness_platform_delegatetoken" "account_level" {
  name       = "account-delegate-token"
  account_id = "account_id"
}

# Look up a delegate token at organization level
data "harness_platform_delegatetoken" "org_level" {
  name       = "org-delegate-token"
  account_id = "account_id"
  org_id     = "org_id"
}

# Look up a delegate token at project level
data "harness_platform_delegatetoken" "project_level" {
  name       = "project-delegate-token"
  account_id = "account_id"
  org_id     = "org_id"
  project_id = "project_id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Account Identifier for the Entity
- `name` (String) Name of the delegate token

### Optional

- `org_id` (String) Org Identifier for the Entity
- `project_id` (String) Project Identifier for the Entity
- `token_status` (String) Status of Delegate Token (ACTIVE or REVOKED). If left empty both active and revoked tokens will be assumed

### Read-Only

- `created_at` (Number) Time when the delegate token is created. This is an epoch timestamp.
- `created_by` (Map of String) created by details
- `id` (String) The ID of this resource.
- `revoke_after` (Number) Epoch time in milliseconds after which the token will be marked as revoked. There can be a delay of up to one hour from the epoch value provided and actual revoking of the token.
- `value` (String) Value of the delegate token. Encoded in base64.
