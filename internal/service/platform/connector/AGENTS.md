# AGENTS.md - Connector Resources

## Purpose

This package implements Terraform resources and data sources for Harness connectors. Connectors provide integration with external systems like cloud providers, monitoring tools, ticketing systems, and secret managers.

## Package Structure

```
connector/
├── connector.go              # Base functions shared by all connectors
├── data_source_connector.go  # Generic connector data source
├── secret_ref_util.go        # Secret reference utilities
├── secretManagers/           # Secret manager connector implementations
│   ├── aws_kms.go
│   ├── aws_secrets_manager.go
│   ├── azure_key_vault.go
│   ├── gcp_kms.go
│   ├── gcp_secret_manager.go
│   ├── hashicorp_vault.go
│   └── ...
├── appdynamics.go            # AppDynamics connector
├── aws_cc.go                 # AWS Cloud Cost connector
├── azure_cloud_cost.go       # Azure Cloud Cost connector
├── datadog.go                # Datadog connector
├── dynatrace.go              # Dynatrace connector
├── elasticsearch.go          # Elasticsearch connector
├── gcp_cloud_cost.go         # GCP Cloud Cost connector
├── jdbc.go                   # JDBC database connector
├── jira.go                   # Jira connector
├── kubernetes_cloud_cost.go  # Kubernetes Cloud Cost connector
├── newrelic.go               # New Relic connector
├── pagerduty.go              # PagerDuty connector
├── prometheus.go             # Prometheus connector
├── service_now.go            # ServiceNow connector
├── splunk.go                 # Splunk connector
├── sumologic.go              # Sumo Logic connector
└── *_test.go                 # Test files for each connector
```

## Base Functions (connector.go)

All connector resources should use these base functions:

```go
// For resource Read operations
resourceConnectorReadBase(ctx, d, meta, connType)

// For data source Read operations
dataConnectorReadBase(ctx, d, meta, connType)

// For Create/Update operations
resourceConnectorCreateOrUpdateBase(ctx, d, meta, connector)

// For Delete operations
resourceConnectorDelete(ctx, d, meta)
```

## Creating a New Connector Resource

1. Create `<connector_type>.go` with:
   - `Resource<ConnectorType>()` function returning `*schema.Resource`
   - Schema definition using `helpers.MergeSchemas`
   - CRUD context functions that call base functions

2. Create `<connector_type>_data_source.go` for the data source

3. Create `<connector_type>_test.go` and `<connector_type>_data_source_test.go` for tests

4. Register in `internal/provider/provider.go`:
   - Add import alias (e.g., `pl_connector_newtype "github.com/harness/terraform-provider-harness/internal/service/platform/connector"`)
   - Add to `ResourcesMap` and `DataSourcesMap`

## Common Patterns

### Schema Definition
```go
func Resource<ConnectorType>() *schema.Resource {
    resource := &schema.Resource{
        Description:   "Resource for creating a <Type> connector.",
        ReadContext:   resource<ConnectorType>Read,
        CreateContext: resource<ConnectorType>CreateOrUpdate,
        UpdateContext: resource<ConnectorType>CreateOrUpdate,
        DeleteContext: resourceConnectorDelete,  // Uses shared delete
        Importer:      helpers.MultiLevelResourceImporter,
        Schema: helpers.MergeSchemas(
            commonConnectorSchema(),       // Common fields
            connectorTypeSpecificSchema(), // Type-specific fields
        ),
    }
    return resource
}
```

### Secret References
Use `secret_ref_text` constant for secret reference fields:
```go
const secret_ref_text = "Reference to a secret for the key. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account` to the expression: account.{identifier}."
```

## Testing

Run tests for a specific connector:
```sh
go test -v ./internal/service/platform/connector/... -run TestAccResourceConnector<Type>
```

Run all connector tests:
```sh
go test -v ./internal/service/platform/connector/... -timeout 120m
```

## API Client

Uses the Platform client from the session:
```go
c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
```

Connector API calls:
- `c.ConnectorsApi.GetConnector()`
- `c.ConnectorsApi.CreateConnector()`
- `c.ConnectorsApi.UpdateConnector()`
- `c.ConnectorsApi.DeleteConnector()`

## DOs

- Use base functions from `connector.go` for CRUD operations
- Use `helpers.MergeSchemas` to combine common and specific schemas
- Use `secret_ref_text` constant for secret reference field descriptions
- Add both resource and data source implementations
- Add comprehensive acceptance tests

## DON'Ts

- Don't duplicate common connector logic - use base functions
- Don't hardcode account/org/project IDs - use schema values
- Don't forget to register new connectors in provider.go
- Don't add local .terraform or terraform.tfstate files to git
