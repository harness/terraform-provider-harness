---
page_title: "harness_chaos_load_test_template_revision Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for one revision of a Harness load test hub template. Creates, reads, and deletes that revision. Load Test Manager does not update a revision in place.
---

# harness_chaos_load_test_template_revision (Resource)

Resource for one revision of a Harness load test hub template.

Creates, reads, and deletes that revision. Load Test Manager does not update a revision in place. `harness_chaos_load_test_template` updates a revision only while that resource's revision is still the latest. `revision` must be a real revision id. `latest` is reserved for syncing a load test. Fields not set are inherited from the previous revision, and that inherited value does not replace the revision. Delete is refused when a REFERENCE load test still pins this revision. Leave `org_id` and `project_id` empty for an account hub.

## Example Usage

```terraform
resource "harness_chaos_load_test_template_revision" "v2" {
  org_id             = "default"
  project_id         = "demo"
  hub_identity       = "hub"
  template_identity  = harness_chaos_load_test_template.checkout.identifier
  revision           = "2"
  name               = "Checkout v2"
}
```

## Import

```shell
# Project hub
terraform import harness_chaos_load_test_template_revision.example <org_id>/<project_id>/<hub_identity>/<template_identity>/<revision>

# Organization hub
terraform import harness_chaos_load_test_template_revision.example <org_id>//<hub_identity>/<template_identity>/<revision>

# Account hub
terraform import harness_chaos_load_test_template_revision.example //<hub_identity>/<template_identity>/<revision>
```
