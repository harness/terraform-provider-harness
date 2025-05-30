---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_har_registry Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating and managing Harness Registries.
---

# harness_platform_har_registry (Resource)

Resource for creating and managing Harness Registries.

## Example Usage

```terraform
# Example of a Virtual Registry
resource "harness_platform_har_registry" "virtual_registry" {
  identifier    = "virtual_docker_registry"
  description   = "Virtual Docker Registry"
  space_ref     = "accountId/orgId/projectId"
  package_type  = "DOCKER"

  config {
    type = "VIRTUAL"
    upstream_proxies = ["registry1", "registry2"]
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of an Upstream Registry with Authentication
resource "harness_platform_har_registry" "upstream_registry" {
  identifier    = "upstream_helm_registry"
  description   = "Upstream Helm Registry"
  space_ref     = "accountId/orgId/projectId"
  package_type  = "HELM"

  config {
    type   = "UPSTREAM"
    source = "CUSTOM"
    url    = "https://helm.sh"
    auth {
      auth_type         = "UserPassword"
      user_name         = "registry_user"
      secret_identifier = "registry_password"
      secret_space_path = "accountId/orgId/projectId"
    }
  }
  parent_ref = "accountId/orgId/projectId"
}
```

## Schema

### Required

- **`identifier`** (String) - Unique identifier of the registry. This field cannot be changed after creation.

  Example:
  ```terraform
  identifier = "my_registry"
  ```

- **`parent_ref`** (String) - The reference to the parent entity (account, organization, or project) under which this registry exists.

  Example:
  ```terraform
  parent_ref = "accountId/orgId/projectId"
  ```

- **`space_ref`** (String) - The reference to the space where the registry is stored.

  Example:
  ```terraform
  space_ref = "accountId/orgId/projectId"
  ```

- **`package_type`** (String) - The type of package supported by the registry. Possible values: `DOCKER`, `HELM`.

  Example:
  ```terraform
  package_type = "DOCKER"
  ```

- **`config`** (List of Object) - Configuration for the registry. (see [below for nested schema](#nestedblock--config)).

  Example:
  ```terraform
  config {
    type = "VIRTUAL"
  }
  ```

### Optional

- **`description`** (String) - Description of the registry.

  Example:
  ```terraform
  description = "Registry for storing Docker images"
  ```

- **`allowedPattern`** (List of String) - List of patterns that are allowed in the registry.

  Example:
  ```terraform
  allowedPattern = ["*/release-*", "*/stable"]
  ```

- **`blockedPattern`** (List of String) - List of patterns that are blocked in the registry.

  Example:
  ```terraform
  blockedPattern = ["*/beta", "*/unstable"]
  ```

- **`url`** (String) - The URL of the registry (Computed value).


### Read-Only

- **`id`** (String) - The unique identifier of the resource generated by Harness.
- **`created_at`** (String) - Timestamp when the registry was created.

<a id="nestedblock--config"></a>
### Nested Schema for `config`

Required:
- **`type`** (String) - Type of registry. Must be either `VIRTUAL` or `UPSTREAM`.

  Example:
  ```terraform
  config {
    type = "VIRTUAL"
  }
  ```

Optional (for VIRTUAL type):
- **`upstream_proxies`** (List of String) - List of upstream proxies. Only applicable when `type` is `VIRTUAL`.

  Example:
  ```terraform
  config {
    type = "VIRTUAL"
    upstream_proxies = ["registry1", "registry2"]
  }
  ```

Required (for UPSTREAM type):
- **`source`** (String) - Source of the upstream registry. Required when `type` is `UPSTREAM`.

  Example:
  ```terraform
  config {
    type = "UPSTREAM"
    source = "docker.io"
  }
  ```

- **`url`** (String) - URL of the upstream registry. Required when `type` is `UPSTREAM` and `package_type` is `HELM`. Must start with http:// or https://.

  Example:
  ```terraform
  config {
    type = "UPSTREAM"
    url = "https://registry.hub.docker.com"
  }
  ```

- **`auth`** (List of Object) - Authentication configuration for upstream registry. (see [below for nested schema](#nestedblock--config--auth))

  Example:
  ```terraform
  config {
    type = "UPSTREAM"
    auth {
      auth_type = "UserPassword"
    }
  }
  ```

- **`auth_type`** (String) - Type of authentication. Must be either `UserPassword` or `Anonymous`.

  Example:
  ```terraform
  config {
    type = "UPSTREAM"
    auth_type = "UserPassword"
  }

  config {
    type = "UPSTREAM"
    auth_type = "Anonymous"
  }
  ```

<a id="nestedblock--config--auth"></a>
### Nested Schema for `config.auth`

Required:
- **`auth_type`** (String) - Type of authentication. Must be either `UserPassword` or `Anonymous`.

  Example:
  ```terraform
  config {
    type = "UPSTREAM"
    auth {
      auth_type = "UserPassword"
    }
  }

  config {
    type = "UPSTREAM"
    auth {
      auth_type = "Anonymous"
    }
  }
  ```

Optional (required if auth_type is UserPassword):
- **`secret_identifier`** (String) - Secret identifier for UserPassword authentication.

- **`secret_space_path`** (String) - Secret space path for UserPassword authentication.

- **`user_name`** (String) - Username for UserPassword authentication.

  Example:
  ```terraform
  config {
    type = "UPSTREAM"
    auth {
      auth_type = "UserPassword"
      secret_identifier = "registry_password"
      secret_space_path = "accountId/orgId/projectId"
      user_name = "registry_user"
    }
  }
  ```