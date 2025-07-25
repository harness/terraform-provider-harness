---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_delegate_approval Resource - terraform-provider-harness"
subcategory: "First Gen"
description: |-
  Resource for approving or rejecting delegates.
---

# harness_delegate_approval (Resource)

Resource for approving or rejecting delegates.

## Example Usage

```terraform
data "harness_delegate" "test" {
  name = "my-delegate"
}

resource "harness_delegate_approval" "test" {
  delegate_id = data.harness_delegate.test.id
  approve     = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `approve` (Boolean) Whether or not to approve the delegate.
- `delegate_id` (String) The id of the delegate.

### Read-Only

- `id` (String) The ID of this resource.
- `status` (String) The status of the delegate.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Import the status of the delegate approval.
terraform import harness_delegate_approval.example <delegate_id>
```
