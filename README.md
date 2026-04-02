# Harness Terraform Provider

[![GitHub release](https://img.shields.io/github/v/release/harness/terraform-provider-harness)](https://github.com/harness/terraform-provider-harness/releases)
[![License](https://img.shields.io/github/license/harness/terraform-provider-harness)](LICENSE.md)
[![Website](https://img.shields.io/badge/website-harness.io-blue)](https://harness.io)

## Overview

The Terraform provider for Harness allows you to manage resources in Harness CD and NextGen platforms using infrastructure as code. This provider enables you to automate the provisioning and management of Harness resources such as applications, pipelines, environments, services, connectors, and more.

## Features

- **First Generation & NextGen Support**: Manage resources in both Harness First Generation and NextGen platforms
- **Comprehensive Resource Management**: Create and manage applications, pipelines, environments, services, connectors, and more
- **GitOps Integration**: Configure GitOps applications, repositories, and clusters
- **Policy Management**: Define and enforce governance policies and policy sets
- **Feature Flag Management**: Configure and manage feature flags for your applications
- **AutoStopping Rules**: Set up cost optimization with AutoStopping rules for various cloud providers

## Documentation

Full, comprehensive documentation is available on the Terraform Registry website:

- [Provider Documentation](https://registry.terraform.io/providers/harness/harness/latest/docs)
- [Harness Terraform Provider Quickstart Guide](https://docs.harness.io/article/7cude5tvzh-harness-terraform-provider)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.21 (for development)

## Usage

### Provider Configuration

```hcl
terraform {
  required_providers {
    harness = {
      source = "harness/harness"
      version = "0.38.5"
    }
  }
}

# Configure the Harness provider for First Gen resources
provider "harness" {
  endpoint   = "https://app.harness.io/gateway"
  account_id = "YOUR_HARNESS_ACCOUNT_ID"
  api_key    = "YOUR_HARNESS_API_KEY"
}

# Configure the Harness provider for Next Gen resources
provider "harness" {
  endpoint         = "https://app.harness.io/gateway"
  account_id       = "YOUR_HARNESS_ACCOUNT_ID"
  platform_api_key = "YOUR_HARNESS_PLATFORM_API_KEY"
}
```

## Installation

### From Terraform Registry (Recommended)

The provider is available on the [Terraform Registry](https://registry.terraform.io/providers/harness/harness/latest) and will be automatically downloaded when you run `terraform init` in a configuration that requires it.

### Building and Testing Locally

1. Clone the repository:
   ```sh
   git clone https://github.com/harness/terraform-provider-harness.git
   cd terraform-provider-harness
   ```

2. Set up your development environment:
   ```sh
   make dev-setup
   ```
   This installs all Go dependencies and development tools.

3. Build and install the provider locally:
   ```sh
   make install
   ```
   This builds the provider and installs it to your local Terraform plugins directory.

4. Configure Terraform to use your local build by adding this to `~/.terraformrc`:
   ```hcl
   provider_installation {
     dev_overrides {
       "registry.terraform.io/harness/harness" = "~/.terraform.d/plugins/registry.terraform.io/harness/harness/0.99.0-dev/<OS>_<ARCH>"
     }
     direct {}
   }
   ```
   Replace `<OS>_<ARCH>` with your platform (e.g., `darwin_arm64`, `linux_amd64`).

5. To remove the local installation:
   ```sh
   make uninstall
   ```

## Development

### Quick Start

```sh
make help        # Show all available commands
make dev-setup   # Set up development environment
make build       # Build the provider
make test        # Run unit tests
make install     # Install provider locally
```

### Available Make Commands

Run `make help` to see all available commands. Here's a summary:

#### Build Commands

| Command | Description |
|---------|-------------|
| `make build` | Build the provider binary |
| `make build-all` | Build for all supported platforms (darwin, linux, windows) |
| `make install` | Build and install to local Terraform plugins directory |
| `make uninstall` | Remove locally installed provider |
| `make clean` | Clean build artifacts and Go caches |
| `make clean-plugins` | Remove all local plugin versions for this provider |

#### Testing Commands

| Command | Description |
|---------|-------------|
| `make test` | Run unit tests |
| `make testacc` | Run acceptance tests (requires `HARNESS_*` env vars) |
| `make test-coverage` | Run tests with coverage report |
| `make sweep` | Run sweepers to clean up test resources (use with caution) |

#### Code Quality Commands

| Command | Description |
|---------|-------------|
| `make fmt` | Format Go source code |
| `make fmt-check` | Check if code is properly formatted |
| `make vet` | Run go vet static analysis |
| `make lint` | Run golangci-lint |
| `make check` | Run all code quality checks |

#### Documentation Commands

| Command | Description |
|---------|-------------|
| `make docs` | Generate provider documentation |
| `make docs-validate` | Validate provider documentation |
| `make changelog` | Generate changelog from `.changelog` entries |

#### Development Commands

| Command | Description |
|---------|-------------|
| `make deps` | Download and tidy Go dependencies |
| `make deps-upgrade` | Upgrade all dependencies |
| `make tools` | Install development tools (tfplugindocs, changelog-build, golangci-lint) |
| `make dev-setup` | Complete development environment setup |
| `make version` | Show version information |

### Running Tests

#### Unit Tests

```sh
make test
```

#### Acceptance Tests

Acceptance tests create real resources in your Harness account. Set the required environment variables first:

```sh
export HARNESS_ACCOUNT_ID="your-account-id"
export HARNESS_API_KEY="your-api-key"
export HARNESS_PLATFORM_API_KEY="your-platform-api-key"

make testacc
```

#### Test Coverage

```sh
make test-coverage
# Opens coverage.html in your browser
```

### Generating Documentation

Provider documentation is auto-generated from schema descriptions:

```sh
make docs
```

Generated docs are placed in the `docs/` directory.

### Generating Changelog

Changelog entries are stored in the `.changelog/` directory. To generate a new changelog:

```sh
make changelog
```

To add a changelog entry, create a file in `.changelog/<PR_NUMBER>.txt`:

```
```release-note:enhancement
resource/harness_platform_connector: Added support for new authentication method
```
```

Supported types: `enhancement`, `bug`, `feature`, `new-resource`, `new-data-source`, `breaking-change`, `note`

### Cleaning Up

```sh
# Clean build artifacts
make clean

# Remove all locally installed plugin versions
make clean-plugins

# Remove test resources (sweepers)
make sweep
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Run tests and linting (`make check && make test`)
4. Commit your changes (`git commit -m 'Add amazing feature'`)
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

## Support

If you have any questions or need assistance:

- Open a [new issue](https://github.com/harness/terraform-provider-harness/issues/new)
- Join our [Slack community](https://harnesscommunity.slack.com/archives/C02G9CUNF1S)
- Visit our [documentation](https://docs.harness.io)

## License

This project is licensed under the terms of the [LICENSE](LICENSE.md) file.
