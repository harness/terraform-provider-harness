# Custom webhook integration. The webhook URL is in endpoints once created.
resource "harness_platform_idp_integration" "orders" {
  identifier   = "orders_webhook"
  name         = "Orders webhook"
  description  = "Imports services pushed by the order system"
  type         = "Custom"
  mode         = "webhook"
  source_label = "orders"
  tags         = ["team:idp", "webhook"]

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

  kind {
    kind        = "custom"
    auto_import = true
    action      = "merge"
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
    kind             = "cmdb_ci_service"
    auto_import      = true
    replication_mode = "incremental"
    filter           = ".operational_status == \"1\""
    # ServiceNow kinds have no default entity kind or type, so set both. They can't change later.
    idp_kind = "component"
    idp_type = "service"
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
