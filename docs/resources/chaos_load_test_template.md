---
page_title: "harness_chaos_load_test_template Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for a Harness load test hub template. Creates a template and updates the revision stored at create. A newer published revision is not overwritten. Delete returns an error when a load test still references the template.
---

# harness_chaos_load_test_template (Resource)

Resource for a Harness load test hub template.

Creates a template and updates the revision id stored at create. A new published revision is a `harness_chaos_load_test_template_revision`. Publishing a newer revision does not change this resource's `revision`, and reads keep returning the stored revision. Applying this resource after that fails instead of overwriting the newer revision. Delete returns an error when a load test still references the template. `revision` on create is a real revision id. `latest` is reserved for syncing a load test. Leave `org_id` and `project_id` empty for an account hub. Set `org_id` and leave `project_id` empty for an organization hub. An empty description, environment, infrastructure, infra type, script, or script source keeps the stored value. The identifier and hub identity allow letters, digits, underscores, hyphens, and `$`.

## Example Usage

```terraform
resource "harness_chaos_load_test_template" "checkout" {
  org_id       = "default"
  project_id   = "demo"
  hub_identity = "hub"
  identifier   = "checkout"
  name         = "Checkout"
  revision     = "1"
  tool_type    = "K6"
  tool_config = jsonencode({
    mode = "script"
    script = {
      content = "a2s="
    }
  })
}
```

## Import

```shell
# Project hub
terraform import harness_chaos_load_test_template.example <org_id>/<project_id>/<hub_identity>/<identifier>

# Organization hub
terraform import harness_chaos_load_test_template.example <org_id>//<hub_identity>/<identifier>

# Account hub
terraform import harness_chaos_load_test_template.example //<hub_identity>/<identifier>
```
