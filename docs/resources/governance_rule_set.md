---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_governance_rule_set Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating, updating, and managing rule.
---

# harness_governance_rule_set (Resource)

Resource for creating, updating, and managing rule.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cloud_provider` (String) The cloud provider for the rule set. It should be either AWS, AZURE or GCP.
- `name` (String) Name of the rule set.
- `rule_ids` (List of String) List of rule IDs

### Optional

- `description` (String) Description for rule set.

### Read-Only

- `id` (String) The ID of this resource.
- `rule_set_id` (String) Id of the rule.
