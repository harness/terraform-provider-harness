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

### Required Arguments

* `identifier` - (String) Unique identifier of the API Key.
* `name` - (String) Name of the API Key.
* `parent_id` - (String) Parent entity identifier of the API Key.
* `apikey_type` - (String) Type of the API Key. Valid value: `USER`.
* `account_id` - (String) Harness account identifier.

### Optional Arguments

* `org_id` - (String) Organization identifier. Required for organization-level and project-level API Keys.
* `project_id` - (String) Project identifier. Required for project-level API Keys.
* `description` - (String) Description of the API Key.
* `tags` - (Set of String) Tags to associate with the API Key.
* `default_time_to_expire_token` - (Number) Default expiration time (in milliseconds) for tokens generated using this API Key.

### Read-Only Attributes

* `id` - (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# Import account-level API Key
terraform import harness_platform_apikey.account_level <parent_id>/<apikey_id>/<apikey_type>

# Import organization-level API Key
terraform import harness_platform_apikey.org_level <org_id>/<parent_id>/<apikey_id>/<apikey_type>

# Import project-level API Key
terraform import harness_platform_apikey.project_level <org_id>/<project_id>/<parent_id>/<apikey_id>/<apikey_type>
