---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_token Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for retrieving a Harness ApiKey Token.
---

# harness_platform_token (Data Source)

Data source for retrieving a Harness ApiKey Token.

## Example Usage

```terraform
data "harness_platform_token" "test" {
  identifier  = "test_token"
  parent_id   = "apikey_parent_id"
  account_id  = "account_id"
  org_id      = "org_id"
  project_id  = "project_id"
  apikey_id   = "apikey_id"
  apikey_type = "USER"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Account Identifier for the Entity
- `apikey_id` (String) Identifier of the API Key
- `apikey_type` (String) Type of the API Key
- `identifier` (String) Unique identifier of the resource.
- `parent_id` (String) Parent Entity Identifier of the API Key

### Optional

- `email` (String) Email Id of the user who created the Token
- `encoded_password` (String) Encoded password of the Token
- `name` (String) Name of the resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `scheduled_expire_time` (Number) Scheduled expiry time in milliseconds
- `username` (String) Name of the user who created the Token
- `valid` (Boolean) Boolean value to indicate if Token is valid or not.
- `valid_from` (Number) This is the time from which the Token is valid. The time is in milliseconds
- `valid_to` (Number) This is the time till which the Token is valid. The time is in milliseconds

### Read-Only

- `description` (String) Description of the resource.
- `id` (String) The ID of this resource.
- `tags` (Set of String) Tags to associate with the resource.
