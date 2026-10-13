# Custom webhook integration. The webhook URL is in endpoints once created.
resource "harness_platform_idp_integration" "orders" {
  identifier  = "orders_webhook"
  name        = "Orders webhook"
  description = "Imports services pushed by the order system"
  type        = "Custom"
  tags        = ["team:idp", "webhook"]

  custom {
    mode         = "webhook"
    source_label = "orders"

    webhook_auth {
      auth_type   = "hmac_auth"
      hash_algo   = "hmac_sha256"
      auth_header = "X-Signature"
      secret_ref  = "account.orders_webhook_secret"
    }

    # Optional; the kind's default columns are used when omitted.
    layout {
      column {
        title = "Service"
        path  = ".spec.service"
      }
    }
  }

  kind {
    kind        = "custom"
    auto_import = true
    action      = "merge"

    # Sent with the create, so the first webhook already uses it. Omitted fields keep the defaults.
    mapping {
      ingestion {
        identifier = ".service.id"
        name       = ".service.name"
        properties = {
          team = ".service.team"
        }
      }
      # Merges payloads into entities that already exist. Required for kinds with action merge.
      correlation {
        source_path = "service.id"
        target_path = "metadata.identifier"
      }
    }
  }
}

# Harness CD integration, reading services from every scope of the account.
resource "harness_platform_idp_integration" "harness_cd" {
  identifier = "harness_cd"
  name       = "Harness CD"
  type       = "HarnessCD"

  scopes {
    all = true
  }

  kind {
    kind        = "service"
    auto_import = true
  }
}

# ServiceNow CMDB integration in an organization, synced every 12 hours.
resource "harness_platform_idp_integration" "servicenow" {
  identifier       = "servicenow_prod"
  name             = "ServiceNow prod"
  org_id           = "default"
  type             = "ServiceNow"
  update_frequency = "12h"

  connector {
    inline {
      credentials_ref = "account.servicenow"
    }
  }

  kind {
    kind        = "cmdb_ci_service"
    auto_import = true
    filter      = ".operational_status == \"1\""
    # ServiceNow kinds have no default entity kind or type, so set both. They can't change later.
    idp_kind = "component"
    idp_type = "service"

    mapping {
      display_name_field = "name"

      # Source fields are top-level fields of the kind's catalog stream.
      field_mapping {
        source_field = "owned_by"
        target_field = "spec.owner"
      }
      field_mapping {
        source_field = "environment"
        target_field = "metadata.environment"
      }
    }
  }
}

# Harness CI integration. Every kind needs a correlation, so each kind has a mapping.
resource "harness_platform_idp_integration" "harness_ci" {
  identifier = "harness_ci"
  name       = "Harness CI"
  type       = "HarnessCI"

  kind {
    kind = "pipeline"

    mapping {
      correlation {
        # A field of the emitted entity data, matched against existing entities.
        source_path = "metadata.integration_properties.HarnessCI.pipelineName"
        target_path = "metadata.name"
      }
    }
  }
}

# GitHub integration for one organization's repositories.
resource "harness_platform_idp_integration" "github" {
  identifier = "github"
  name       = "GitHub"
  type       = "GitHub"

  connector {
    inline {
      credentials_ref = "account.github"
    }
  }

  github {
    orgs = ["acme"]
  }

  kind {
    kind = "repo_details"
  }
}
