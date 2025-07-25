---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_cluster_orchestrator Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating ClusterOrchestrators.
---

# harness_cluster_orchestrator (Resource)

Resource for creating ClusterOrchestrators.

## Example Usage

```terraform
resource "harness_cluster_orchestrator" "test" {
  name             = "name"
  cluster_endpoint = "http://test.test.com"
  k8s_connector_id = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_endpoint` (String) Endpoint of the k8s cluster being onboarded under the orchestrator
- `k8s_connector_id` (String) ID of the Harness Kubernetes Connector Being used
- `name` (String) Name of the Orchestrator

### Read-Only

- `id` (String) The ID of this resource.
