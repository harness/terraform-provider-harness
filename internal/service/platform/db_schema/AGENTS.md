# AGENTS.md - DB Schema Resource

## Purpose

This package implements the Terraform resource and data source for Harness DBDevOps Schema. A database schema represents the configuration for managing database migrations using tools like Liquibase or Flyway.

## Key Files

- `resource_db_schema.go` - Main resource implementation with CRUD operations
- `resource_db_schema_test.go` - Acceptance tests for the resource
- `data_source_db_schema.go` - Data source implementation for reading existing schemas
- `data_source_db_schema_test.go` - Acceptance tests for the data source

## Resource Schema

The DB Schema resource supports:
- **service** - The service associated with the schema
- **type** - Schema type: `SCRIPT` or `REPOSITORY`
- **migration_type** - Migration tool: `Liquibase` or `Flyway`
- **use_percona** - Enable percona-toolkit for the database schema
- **schema_source** - Repository-based schema configuration (connector, location, repo)
- **changelog_script** - Script-based schema configuration (image, shell, etc.)

Note: `schema_source` and `changelog_script` are mutually exclusive.

## API Client

Uses the DBOps client from the session:
```go
c, ctx := meta.(*internal.Session).GetDBOpsClientWithContext(ctx)
```

## Error Handling

Use helpers specific to DBOps:
- `helpers.HandleDBOpsApiError(err, d, httpResp)` - For create/update/delete
- `helpers.HandleDBOpsReadApiError(err, d, httpResp)` - For read operations

## Testing

Run tests for this package:
```sh
go test -v ./internal/service/platform/db_schema/... -timeout 120m
```

Run acceptance tests (requires HARNESS_* env vars):
```sh
TF_ACC=1 go test -v ./internal/service/platform/db_schema/... -timeout 120m
```

## Dependencies

- `github.com/harness/harness-go-sdk/harness/dbops` - DBOps API client
- `github.com/harness/terraform-provider-harness/helpers` - Shared helpers and error handling

## Project-Level Fields

Uses `helpers.SetProjectLevelResourceSchema` to add common fields:
- `identifier`, `name`, `description`
- `org_id`, `project_id`
- `tags`

## DOs

- Use `helpers.HandleDBOpsApiError` for error handling (not `HandleApiError`)
- Follow the existing pattern for schema_source vs changelog_script mutual exclusivity
- Add acceptance tests for new functionality
- Update documentation schema descriptions

## DON'Ts

- Don't use `HandleApiError` - use `HandleDBOpsApiError` for DBOps resources
- Don't add local .terraform or terraform.tfstate files to git
