# AGENTS.md

## Project Overview

Terraform provider for Harness that enables infrastructure-as-code management of Harness CD and NextGen platform resources. This provider allows automation of applications, pipelines, environments, services, connectors, feature flags, GitOps configurations, and AutoStopping rules.

## Build System

- **Language**: Go 1.24
- **Build tool**: Make
- **Build command**: `make build`
- **Install locally**: `make install`
- **Clean**: `make clean`

## Testing

- **Run all unit tests**: `make test`
- **Run acceptance tests**: `make testacc` (requires HARNESS_* environment variables)
- **Run tests with coverage**: `make test-coverage`
- **Run single test file**: `go test -v ./internal/service/platform/connector/... -run TestAccResourceConnector`
- **Run single test case**: `go test -v ./path/to/package -run TestFunctionName`

### Acceptance Tests

Acceptance tests create real resources in Harness. Required environment variables:
- `HARNESS_ACCOUNT_ID`
- `HARNESS_API_KEY`
- `HARNESS_PLATFORM_API_KEY`

## Linting & Formatting

- **Format code**: `make fmt`
- **Check formatting**: `make fmt-check`
- **Run vet**: `make vet`
- **Run linter**: `make lint` (uses golangci-lint)
- **Run all checks**: `make check` (fmt-check + vet)

## Git Workflow

- **Commit format**: `type: [JIRA-TICKET]: Description`
  - Types: `feat`, `fix`, `chore`, `refactor`, `docs`, `test`
  - Jira prefixes: DBOPS, CDS, CCM, IAC, PL, AH, and others
  - Example: `feat: [DBOPS-2019]: Add usePercona field support in schema resource`
- **Default branch**: `main`

## DOs

- Always run `make check` before committing to verify formatting and static analysis
- Add changelog entries in `.changelog/` for user-facing changes (see README for format)
- Run `make test` to ensure unit tests pass before pushing
- Follow existing code patterns in the codebase
- Use descriptive commit messages with Jira ticket references
- Run `make docs` after modifying resource schemas to regenerate documentation
- Use `make dev-setup` when setting up the development environment for the first time

## DON'Ts

- Never force push to main/master
- Never commit secrets, .env files, or credentials
- Never run `make sweep` without explicit confirmation - it destroys test resources and can be dangerous if misconfigured
- Never skip code quality checks before committing
- Never modify generated documentation files in `docs/` directly - regenerate with `make docs`

## Commands to Never Run

- `git push --force origin main`
- `git push --force origin master`
- `git commit --no-verify` or `git push --no-verify`
- `make sweep` (without explicit user confirmation)
- `rm -rf /` or any destructive recursive deletes

## Project Structure

```
terraform-provider-harness/
├── main.go                 # Provider entry point
├── Makefile                # Build, test, and development commands
├── go.mod                  # Go module definition
├── internal/               # Internal packages (not importable externally)
│   ├── provider/           # Terraform provider configuration
│   ├── service/            # Resource implementations by service
│   │   ├── cd/             # First Generation CD resources
│   │   ├── cd_nextgen/     # NextGen CD resources
│   │   ├── chaos/          # Chaos Engineering resources
│   │   ├── pipeline/       # Pipeline resources
│   │   ├── platform/       # Platform resources (connectors, secrets, etc.)
│   │   └── service_discovery/  # Service discovery resources
│   ├── acctest/            # Acceptance test utilities
│   ├── sweep/              # Resource sweepers for test cleanup
│   ├── test/               # Test utilities
│   └── utils/              # Utility functions
├── docs/                   # Generated Terraform documentation
├── examples/               # Example Terraform configurations
├── templates/              # Documentation templates
├── scripts/                # Build and utility scripts
├── helpers/                # Helper utilities
└── .changelog/             # Changelog entries for releases
```

## Important Packages

These are the most actively developed areas based on recent commit history:

- `internal/service/platform/` - Platform resources (connectors, secrets, users, etc.)
- `internal/service/cd_nextgen/` - NextGen CD resources (environments, services, infrastructure)
- `internal/service/pipeline/` - Pipeline resources
- `internal/acctest/` - Acceptance test utilities

## Resource Development Guidelines

When creating or modifying Terraform resources:

1. **Schema Definition**: Define resource schema in the resource file with clear descriptions
2. **CRUD Operations**: Implement Create, Read, Update, Delete functions
3. **Testing**: Write both unit tests and acceptance tests
4. **Documentation**: Schema descriptions auto-generate docs via `make docs`
5. **Changelog**: Add entry in `.changelog/<PR_NUMBER>.txt`

## Dependencies

Key dependencies (from go.mod):
- `github.com/harness/harness-go-sdk` - Harness Go SDK
- `github.com/harness/harness-openapi-go-client` - Harness OpenAPI client
- `github.com/hashicorp/terraform-plugin-sdk/v2` - Terraform Plugin SDK
- `github.com/stretchr/testify` - Testing utilities

## Language-Specific Guidelines

This is a Go codebase. Follow Go conventions:
- Use `gofmt` for formatting (handled by `make fmt`)
- Follow effective Go guidelines
- Use meaningful variable and function names
- Handle errors explicitly
- Write table-driven tests where appropriate
