# AGENTS.md - DB Instance Resource

## Purpose

This package implements the Terraform resource and data source for Harness DBDevOps Instance. A database instance represents a specific database that is managed by a parent database schema. It connects a database (via a connector) to a schema definition.

## Key Files

- `resource_db_instance.go` - Main resource implementation with CRUD operations
- `resource_db_instance_test.go` - Acceptance tests for the resource
- `data_source_db_instance.go` - Data source implementation for reading existing instances
- `data_source_db_instance_test.go` - Acceptance tests for the data source

## Resource Schema

The DB Instance resource supports:
- **schema** (required) - The identifier of the parent database schema
- **connector** (required) - The connector to the database
- **branch** - The branch of changeSet repository
- **context** - The Liquibase context
- **substitute_properties** - Properties to substitute in changelog migration script (map)

## Relationship to DB Schema

A DB Instance belongs to a DB Schema. The parent schema is specified by the `schema` field. API calls include the schema identifier in the path:
```go
c.DatabaseInstanceApi.V1GetProjDbSchemaInstance(ctx, orgId, projectId, schemaId, instanceId, &opts)
```

## API Client

Uses the DBOps client from the session:
```go
c, ctx := meta.(*internal.Session).GetDBOpsClientWithContext(ctx)
```

## Error Handling

Use helpers specific to DBOps:
- `helpers.HandleDBOpsApiError(err, d, httpResp)` - For create/update/delete
- `helpers.HandleDBOpsReadApiError(err, d, httpResp)` - For read operations

## Resource Importer

Uses a special importer for DB instances:
```go
Importer: helpers.DBInstanceResourceImporter
```

## Testing

Run tests for this package:
```sh
go test -v ./internal/service/platform/db_instance/... -timeout 120m
```

Run acceptance tests (requires HARNESS_* env vars):
```sh
TF_ACC=1 go test -v ./internal/service/platform/db_instance/... -timeout 120m
```

## Dependencies

- `github.com/harness/harness-go-sdk/harness/dbops` - DBOps API client
- `github.com/harness/terraform-provider-harness/helpers` - Shared helpers and error handling
- Depends on `db_schema` resource existing first

## Project-Level Fields

Uses `helpers.SetProjectLevelResourceSchema` to add common fields:
- `identifier`, `name`, `description`
- `org_id`, `project_id`
- `tags`

## DOs

- Use `helpers.HandleDBOpsApiError` for error handling (not `HandleApiError`)
- Always reference the parent schema via the `schema` field
- Add acceptance tests for new functionality
- Update documentation schema descriptions

## DON'Ts

- Don't use `HandleApiError` - use `HandleDBOpsApiError` for DBOps resources
- Don't forget the parent schema reference in API calls
- Don't add local .terraform or terraform.tfstate files to git
