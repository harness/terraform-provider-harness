# FME (Feature Management Engine) Integration

This document describes the integration of Split.io's Feature Management Engine (FME) into the Harness Terraform provider.

## Overview

The FME integration allows you to manage Split.io feature flags, environments, and API keys through Terraform using the existing Harness provider authentication.

## Resources

### `harness_fme_environment`

Manages Split.io environments.

```hcl
resource "harness_fme_environment" "dev" {
  name       = "development"
  production = false
}

resource "harness_fme_environment" "prod" {
  name       = "production"
  production = true
}
```

**Arguments:**
- `name` (Required) - Name of the environment
- `production` (Optional) - Whether this is a production environment (default: false)

**Attributes:**
- `id` - Unique identifier of the environment

### `harness_fme_api_key`

Manages Split.io API keys for environments.

```hcl
resource "harness_fme_api_key" "client_key" {
  environment_id = harness_fme_environment.dev.id
  name          = "My Client Side Key"
  type          = "client_side"
}

resource "harness_fme_api_key" "server_key" {
  environment_id = harness_fme_environment.prod.id
  name          = "My Server Side Key"
  type          = "server_side"
}
```

**Arguments:**
- `environment_id` (Required) - ID of the environment this API key belongs to
- `name` (Required) - Name of the API key
- `type` (Required) - Type of API key. Must be either `client_side` or `server_side`

**Attributes:**
- `id` - Unique identifier of the API key
- `key` - The actual API key value (sensitive)

### `harness_fme_split`

Manages Split.io feature flags (splits).

```hcl
resource "harness_fme_split" "new_checkout_flow" {
  workspace_id = data.harness_fme_workspace.main.id
  name         = "new-checkout-flow"
  description  = "Enable new checkout flow for better conversion"
}
```

**Arguments:**
- `workspace_id` (Required) - ID of the workspace this split belongs to
- `name` (Required) - Name of the split
- `description` (Required) - Description of the split

**Attributes:**
- `id` - Unique identifier of the split
- `creation_time` - Creation time of the split
- `rollout_status_timestamp` - Rollout status timestamp

## Data Sources

### `harness_fme_workspace`

Retrieves workspace information.

```hcl
data "harness_fme_workspace" "main" {
  id = "ws_123456789"
}
```

**Arguments:**
- `id` (Required) - Unique identifier of the workspace

**Attributes:**
- `name` - Name of the workspace
- `type` - Type of the workspace
- `display_name` - Display name of the workspace

## Authentication

The FME integration uses the existing Harness provider authentication but sends API calls to Split.io:

```hcl
provider "harness" {
  endpoint           = "https://app.harness.io/gateway"
  account_id         = "your-account-id"
  platform_api_key   = "your-split-api-key"  # Split.io API key
}
```

## API Details

- **Base URL**: `https://api.split.io/internal/api/v2` (most resources)
- **Flag Sets URL**: `https://api.split.io/api/v3/flag-sets` (flag sets use different API version)
- **Authentication**: `X-API-Key` header
- **Rate Limiting**: 2-second retry on 429 responses

## Complete Example

```hcl
# Provider configuration
provider "harness" {
  endpoint           = "https://app.harness.io/gateway"
  account_id         = "your-account-id"
  platform_api_key   = "your-split-api-key"
}

# Get workspace
data "harness_fme_workspace" "main" {
  id = "ws_123456789"
}

# Create environments
resource "harness_fme_environment" "staging" {
  name       = "staging"
  production = false
}

resource "harness_fme_environment" "production" {
  name       = "production"
  production = true
}

# Create API keys
resource "harness_fme_api_key" "staging_client" {
  environment_id = harness_fme_environment.staging.id
  name          = "Staging Client Key"
  type          = "client_side"
}

resource "harness_fme_api_key" "prod_server" {
  environment_id = harness_fme_environment.production.id
  name          = "Production Server Key"
  type          = "server_side"
}

# Create feature flags
resource "harness_fme_split" "user_onboarding_v2" {
  workspace_id = data.harness_fme_workspace.main.id
  name         = "user-onboarding-v2"
  description  = "New user onboarding flow with improved UX"
}

# Outputs
output "workspace_info" {
  value = {
    name         = data.harness_fme_workspace.main.name
    display_name = data.harness_fme_workspace.main.display_name
  }
}

output "api_keys" {
  value = {
    staging_client = harness_fme_api_key.staging_client.key
    prod_server    = harness_fme_api_key.prod_server.key
  }
  sensitive = true
}
```