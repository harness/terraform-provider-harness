---
page_title: "harness_platform_apikey Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating and managing Harness API Keys.
---

# harness_platform_apikey (Resource)

Resource for creating and managing Harness API Keys. API Keys can be created at the account, organization, or project level.

## Example Usage

```terraform
# Create API Key at account level
resource "harness_platform_apikey" "account_level" {
  identifier   = "test_apikey"
  name         = "test_apikey"
  parent_id    = "parent_id"
  apikey_type  = "USER"
  account_id   = "account_id"
}

# Create API Key at organization level
resource "harness_platform_apikey" "org_level" {
  identifier   = "test_apikey"
  name         = "test_apikey"
  parent_id    = "parent_id"
  apikey_type  = "USER"
  account_id   = "account_id"
  org_id       = "org_id"
}

# Create API Key at project level
resource "harness_platform_apikey" "project_level" {
  identifier   = "test_apikey"
  name         = "test_apikey"
  parent_id    = "parent_id"
  apikey_type  = "USER"
  account_id   = "account_id"
  org_id       = "org_id"
  project_id   = "project_id"
}
```

## Schema

### Required

- `account_id` (String) Account Identifier for the Entity
- `apikey_type` (String) Type of the API Key
- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `parent_id` (String) Parent Entity Identifier of the API Key

### Optional

- `description` (String) Description of the resource.
- `tags` (Set of String) Tags to associate with the resource.
- `default_time_to_expire_token` (Number) Default expiration time of the Token within API Key
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# Import account level apikey
terraform import harness_platform_apikey <parent_id>/<apikey_id>/<apikey_type>

# Import org level apikey
terraform import harness_platform_apikey <org_id>/<parent_id>/<apikey_id>/<apikey_type>

# Import project level apikey
terraform import harness_platform_apikey <org_id>/<project_id>/<parent_id>/<apikey_id>/<apikey_type>
```
