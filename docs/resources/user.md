---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_user Resource - terraform-provider-harness"
subcategory: "First Gen"
description: |-
  Resource for creating a Harness user
---

# harness_user (Resource)

Resource for creating a Harness user

## Example Usage

```terraform
resource "harness_user" "john_doe" {
  name  = "John Doe"
  email = "john.doe@example.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `email` (String) The email of the user.
- `name` (String) The name of the user.

### Optional

- `group_ids` (Set of String) The groups the user belongs to. This is only used during the creation of the user. The groups are not updated after the user is created. When using this option you should also set `lifecycle = { ignore_changes = ["group_ids"] }`.

### Read-Only

- `id` (String) Unique identifier of the user.
- `is_email_verified` (Boolean) Flag indicating whether or not the users email has been verified.
- `is_imported_from_identity_provider` (Boolean) Flag indicating whether or not the user was imported from an identity provider.
- `is_password_expired` (Boolean) Flag indicating whether or not the users password has expired.
- `is_two_factor_auth_enabled` (Boolean) Flag indicating whether or not two-factor authentication is enabled for the user.
- `is_user_locked` (Boolean) Flag indicating whether or not the user is locked out.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Import using the email address of the user
terraform import harness_user.john_doe john.doe@example.com
```
