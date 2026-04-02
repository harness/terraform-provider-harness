# Chaos Hub V2

This package provides Terraform resources and data sources for managing Harness Chaos Hub using the REST API (V2).

## Overview

Chaos Hub V2 is a REST API-based implementation for managing chaos hubs in Harness. It provides a modern alternative to the GraphQL-based chaos_hub implementation, following the same patterns as other V2 chaos resources like `infrastructure_v2`.

## Files

- `resource_chaos_hub_v2_schema.go` - Schema definitions for the resource
- `resource_chaos_hub_v2.go` - Main resource implementation with CRUD operations
- `data_source_chaos_hub_v2.go` - Data source for querying chaos hubs
- `resource_chaos_hub_v2_test.go` - Resource acceptance tests
- `data_source_chaos_hub_v2_test.go` - Data source acceptance tests

## Resource: harness_chaos_hub_v2

### Example Usage

```hcl
resource "harness_chaos_hub_v2" "example" {
  org_id        = "my_org"
  project_id    = "my_project"
  identity      = "custom-chaos-hub"
  name          = "Custom Chaos Hub"
  connector_ref = "project.github_connector"
  
  repo_branch = "main"
  repo_name   = "chaos-experiments"
  description = "Custom chaos experiments repository"
  
  tags = ["production", "critical"]
}
```

### Schema

#### Required Arguments

- `org_id` - (String, ForceNew) The ID of the organization.
- `project_id` - (String, ForceNew) The ID of the project.
- `identity` - (String, ForceNew) Unique identifier for the chaos hub.
- `name` - (String) Name of the chaos hub.
- `connector_ref` - (String, ForceNew) Reference to the Git connector (format: `scope.connectorId`).

#### Optional Arguments

- `repo_branch` - (String, ForceNew) Git repository branch.
- `repo_name` - (String, ForceNew) Name of the Git repository.
- `description` - (String) Description of the chaos hub.
- `tags` - (List of String) Tags to associate with the chaos hub.

#### Computed Attributes

- `hub_id` - Internal hub ID returned by the API.
- `account_id` - Account ID.
- `is_default` - Whether this is the default chaos hub.
- `is_removed` - Whether the chaos hub has been removed.
- `created_at` - Creation timestamp (Unix epoch).
- `updated_at` - Last update timestamp (Unix epoch).
- `created_by` - User who created the chaos hub.
- `updated_by` - User who last updated the chaos hub.
- `last_synced_at` - Timestamp of the last sync (Unix epoch).
- `repo_url` - Git repository URL.
- `connector_id` - Connector ID (deprecated, use connector_ref).
- `action_template_count` - Number of action templates in the hub.
- `experiment_template_count` - Number of experiment templates in the hub.
- `fault_template_count` - Number of fault templates in the hub.
- `probe_template_count` - Number of probe templates in the hub.

### Import

Chaos hubs can be imported using the format: `org_id/project_id/identity`

```bash
terraform import harness_chaos_hub_v2.example my_org/my_project/custom-chaos-hub
```

## Data Source: harness_chaos_hub_v2

### Example Usage

```hcl
# Query by identity
data "harness_chaos_hub_v2" "by_identity" {
  org_id     = "my_org"
  project_id = "my_project"
  identity   = "custom-chaos-hub"
}

# Query by name
data "harness_chaos_hub_v2" "by_name" {
  org_id     = "my_org"
  project_id = "my_project"
  name       = "Custom Chaos Hub"
}
```

### Schema

#### Required Arguments

- `org_id` - (String) The ID of the organization.
- `project_id` - (String) The ID of the project.

#### Optional Arguments (Exactly One Required)

- `identity` - (String) Unique identifier of the chaos hub.
- `name` - (String) Name of the chaos hub.

#### Computed Attributes

All resource attributes are available as computed attributes in the data source.

## API Endpoints

This implementation uses the following REST API endpoints:

- **POST** `/rest/hubs` - Create chaos hub
- **GET** `/rest/hubs/{hubIdentity}` - Get chaos hub
- **PUT** `/rest/hubs/{hubIdentity}` - Update chaos hub
- **DELETE** `/rest/hubs/{hubIdentity}` - Delete chaos hub
- **GET** `/rest/hubs` - List chaos hubs (for data source name lookup)

## Implementation Details

### Consistency with infrastructure_v2

This implementation follows the exact same patterns as `infrastructure_v2`:

1. **File Structure**: Separate files for schema, resource, data source, and tests
2. **Client Usage**: Uses `GetChaosClientWithContext()` from session
3. **API Calls**: Direct usage of `DefaultApi` methods
4. **Error Handling**: Uses `helpers.HandleApiError()` for consistent error handling
5. **Import Format**: Follows `org_id/project_id/resource_id` pattern
6. **Test Structure**: Mirrors infrastructure_v2 test patterns

### Key Differences from GraphQL Implementation

| Aspect | GraphQL (chaos_hub) | REST (chaos_hub_v2) |
|--------|---------------------|---------------------|
| Client | `chaos.NewChaosHubClient(c)` | `c.DefaultApi` |
| Identifier | Internal `hubID` | User-defined `identity` |
| Connector | Separate `connector_id` + `connector_scope` | Combined `connector_ref` |
| Package | `chaos_hub` | `chaos_hub_v2` |
| Resource Name | `harness_chaos_hub` | `harness_chaos_hub_v2` |

## Testing

### Running Tests

```bash
# Run all tests
go test -v ./internal/service/chaos/chaos_hub_v2/

# Run specific test
go test -v ./internal/service/chaos/chaos_hub_v2/ -run TestAccResourceChaosHubV2_basic

# Run with acceptance test flag
TF_ACC=1 go test -v ./internal/service/chaos/chaos_hub_v2/
```

### Test Coverage

- ✅ Basic resource creation
- ✅ Resource updates
- ✅ Resource import
- ✅ Data source lookup by identity
- ✅ Data source lookup by name
- ✅ Destroy verification

## Best Practices

1. **Use Meaningful Identities**: Choose descriptive, unique identities for your hubs
2. **Scope Connectors Appropriately**: Use project-level connectors for project-specific hubs
3. **Tag Your Hubs**: Use tags for organization and filtering
4. **Document Descriptions**: Provide clear descriptions for team collaboration
5. **Plan ForceNew Changes**: Be aware that changing `identity`, `connector_ref`, `repo_branch`, or `repo_name` requires recreation

## Migration from chaos_hub

If migrating from the GraphQL-based `harness_chaos_hub`:

1. Change resource type to `harness_chaos_hub_v2`
2. Add explicit `identity` field
3. Convert `connector_id` + `connector_scope` to `connector_ref` format
4. Update import statements to new format

Example:

```hcl
# Before (GraphQL)
resource "harness_chaos_hub" "example" {
  name           = "My Hub"
  connector_id   = "my_connector"
  connector_scope = "PROJECT"
  repo_branch    = "main"
}

# After (REST V2)
resource "harness_chaos_hub_v2" "example" {
  org_id        = "my_org"
  project_id    = "my_project"
  identity      = "my-hub"
  name          = "My Hub"
  connector_ref = "project.my_connector"
  repo_branch   = "main"
}
```

## Troubleshooting

### Common Issues

1. **"Hub identity already exists"**
   - Solution: Choose a unique identity or import the existing hub

2. **"Connector not found"**
   - Solution: Verify the `connector_ref` format and that the connector exists

3. **"Permission denied"**
   - Solution: Ensure the API key has chaos module permissions

4. **"Invalid repository"**
   - Solution: Verify the repository exists and is accessible via the connector

## Contributing

When contributing to this implementation:

1. Follow the existing code patterns from `infrastructure_v2`
2. Add tests for new functionality
3. Update documentation
4. Ensure backward compatibility where possible
5. Run tests before submitting

## References

- [Harness Chaos Hub Documentation](https://developer.harness.io/docs/chaos-engineering/chaos-hubs/)
- [Infrastructure V2 Implementation](../infrastructure_v2/)
- [Harness Go SDK](https://github.com/harness/harness-go-sdk)
