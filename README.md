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
- [Go](https://golang.org/doc/install) >= 1.17 (for development)

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

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Build the provider:
   ```sh
   go build -o terraform-provider-harness
   ```

4. Create a file called `local.sh` in the root directory with the following content:
   ```sh
   #!/bin/sh

   version=0.40.2 # Specify in this format
   source=registry.terraform.io/harness/harness
   platform=darwin_amd64 # Use darwin_arm64 for Apple Silicon based Mac

   mkdir -p ~/.terraform.d/plugins/$source/$version/$platform/

   cp terraform-provider-harness ~/.terraform.d/plugins/$source/$version/$platform/terraform-provider-harness
   ```

5. Make the script executable and run it:
   ```sh
   chmod +x local.sh
   ./local.sh
   ```

### Configure Terraform to Use Local Build

1. Update the Terraform CLI configuration file (`.terraformrc` on macOS/Linux, `terraform.rc` on Windows):
   ```hcl
   provider_installation {
     dev_overrides {
       "registry.terraform.io/harness/harness" = "/path/to/terraform-provider-harness"
     }
     direct {}
   }
   ```
**Note**: Ensure the terraform provider version in your configuration matches the version in the `local.sh` script.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

If you have any questions or need assistance:

- Open a [new issue](https://github.com/harness/terraform-provider-harness/issues/new)
- Join our [Slack community](https://harnesscommunity.slack.com/archives/C02G9CUNF1S)
- Visit our [documentation](https://docs.harness.io)

## License

This project is licensed under the terms of the [LICENSE](LICENSE.md) file.
