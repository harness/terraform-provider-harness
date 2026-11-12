---
page_title: "harness_chaos_load_test_template_revision Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for one revision of a Harness load test hub template.
---

# harness_chaos_load_test_template_revision (Data Source)

Data source for one revision of a Harness load test hub template. Leave `org_id` and `project_id` empty for an account hub, or set only `org_id` for an organization hub.

## Example Usage

```terraform
data "harness_chaos_load_test_template_revision" "v2" {
  org_id             = "default"
  project_id         = "demo"
  hub_identity       = "hub"
  template_identity  = "checkout"
  revision           = "2"
}
```
