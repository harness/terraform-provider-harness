---
page_title: "Harness: harness_platform_gitops_filters"
description: |-
  Data source for retrieving a Harness GitOps Filter.
---

# harness_platform_gitops_filters

Data source for retrieving a Harness GitOps Filter.

## Example Usage

```terraform
data "harness_platform_gitops_filters" "example" {
  identifier = "example_filter"
  org_id     = "your_org_id"
  project_id = "your_project_id"
  type       = "APPLICATION"
}

output "filter_properties" {
  value = jsondecode(data.harness_platform_gitops_filters.example.filter_properties)
}
```

## Argument Reference

* `identifier` - (Required) Unique identifier of the GitOps filter to retrieve.
* `type` - (Required) Type of GitOps filter. Currently, only "APPLICATION" is supported.
* `org_id` - (Optional) Organization identifier for the GitOps filter.
* `project_id` - (Optional) Project identifier for the GitOps filter.
* `filter_visibility` - (Optional) Use this to filter by visibility. Valid values are "EveryOne" and "OnlyCreator".

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Name of the GitOps filter.
* `filter_properties` - Properties of the filter entity defined in Harness as a JSON string. This contains filter criteria such as health status, sync status, agent identifiers, clusters, namespaces, and repositories.

### Filter Properties 

The `filter_properties` JSON string may contain the following filter types:

> **Note:** The following filter properties are only valid for filter type "APPLICATION". Different filter types may support different properties.

* `agentIdentifiers` - Array of GitOps agent identifiers.
* `clusters` - Array of cluster URLs.
* `healthStatus` - Array of health status values. May include: "Unknown", "Progressing", "Suspended", "Healthy", "Degraded", "Missing".
* `namespaces` - Array of Kubernetes namespaces.
* `repositories` - Array of Git repository URLs.
* `syncStatus` - Array of sync status values. May include: "OutOfSync", "Synced", "Unknown".

## Usage with JSON Decode

Since the `filter_properties` attribute is returned as a JSON string, you might want to decode it for further processing:

```terraform
locals {
  filter_data = jsondecode(data.harness_platform_gitops_filters.example.filter_properties)
  
  # Access specific filter properties
  health_statuses = local.filter_data.healthStatus
  sync_statuses   = local.filter_data.syncStatus
  namespaces      = local.filter_data.namespaces
}
```
