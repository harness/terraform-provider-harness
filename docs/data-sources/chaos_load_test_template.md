---
page_title: "harness_chaos_load_test_template Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for a Harness load test hub template. Reads the latest revision.
---

# harness_chaos_load_test_template (Data Source)

Data source for a Harness load test hub template. Reads the latest revision. Leave `org_id` and `project_id` empty for an account hub, or set only `org_id` for an organization hub.

## Example Usage

```terraform
data "harness_chaos_load_test_template" "checkout" {
  org_id       = "default"
  project_id   = "demo"
  hub_identity = "hub"
  identifier   = "checkout"
}
```
