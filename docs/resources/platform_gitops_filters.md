---
page_title: "Harness: harness_platform_gitops_filters"
description: |-
  Resource for creating Harness GitOps Filters.
---

# harness_platform_gitops_filters

Resource for creating and managing Harness GitOps Filters.

## Example Usage

```terraform
resource "harness_platform_gitops_filters" "example" {
  name             = "example_filter"
  org_id           = "your_org_id"
  project_id       = "your_project_id"
  identifier       = "example_filter"
  type             = "APPLICATION"
  filter_properties = jsonencode({
    "agentIdentifiers": [
      "your_agent_identifier"
    ],
    "clusters": [
      "https://your-cluster-url"
    ],
    "healthStatus": [
      "Unknown",
      "Progressing",
      "Suspended",
      "Healthy",
      "Degraded",
      "Missing"
    ],
    "namespaces": [
      "your-namespace"
    ],
    "repositories": [
      "your-repo"
    ],
    "syncStatus": [
      "OutOfSync",
      "Synced",
      "Unknown"
    ]
  })
  filter_visibility = "OnlyCreator"
}
```

## Argument Reference

* `identifier` - (Required) Unique identifier of the GitOps filter.
* `name` - (Required) Name of the GitOps filter.
* `type` - (Required) Type of GitOps filter. Currently, only "APPLICATION" is supported.
* `org_id` - (Required) Organization identifier for the GitOps filter.
* `project_id` - (Required) Project identifier for the GitOps filter.
* `filter_properties` - (Required) Properties of the filter entity defined in Harness as a JSON string. All values should be arrays of strings. Example: `jsonencode({"healthStatus": ["Healthy", "Degraded"], "syncStatus": ["Synced"]})`.
* `filter_visibility` - (Optional) Visibility of the filter. Valid values are "EveryOne" and "OnlyCreator". Default is "EveryOne".

## Filter Properties Reference

The `filter_properties` field supports the following filter types:

> **Note:** The following filter properties are only valid for filter type "APPLICATION". Different filter types may support different properties.

* `agentIdentifiers` - Array of GitOps agent identifiers to filter by.
* `clusters` - Array of cluster URLs to filter by.
* `healthStatus` - Array of health status values to filter by. Valid values include: "Unknown", "Progressing", "Suspended", "Healthy", "Degraded", "Missing".
* `namespaces` - Array of Kubernetes namespaces to filter by.
* `repositories` - Array of Git repository URLs to filter by.
* `syncStatus` - Array of sync status values to filter by. Valid values include: "OutOfSync", "Synced", "Unknown".

## Import

GitOps filters can be imported using a composite ID formed of organization ID, project ID, filter ID, and filter type.

```bash
# Format: <org_id>/<project_id>/<filter_id>/<filter_type>
terraform import harness_platform_gitops_filters.example org_id/project_id/filter_id/APPLICATION
```

## Schema Attributes Reference

* `id` - Identifier of the GitOps filter.
