# End-to-end example for the IDP integrations resources and data sources.
# See README.md in this directory for how to run it.

terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

provider "harness" {}

variable "prefix" {
  description = "Prefix for every identifier, so several people can run this in one account."
  type        = string
  default     = "tf_e2e"
}

variable "org_id" {
  description = "Organization for the org-scoped integrations."
  type        = string
  default     = "default"
}

variable "project_id" {
  description = "Project for the HarnessK8s integration; empty for org scope."
  type        = string
  default     = ""
}

variable "webhook_secret_ref" {
  description = "HMAC secret, e.g. account.orders_webhook_secret. When set, the Custom integration uses webhook mode instead of api mode."
  type        = string
  default     = ""
}

variable "k8s_agent_ref" {
  description = "Service discovery agent, e.g. account.my_agent. The HarnessK8s integration is skipped when empty."
  type        = string
  default     = ""
}

variable "k8s_environment" {
  description = "Environment the service discovery agent runs in."
  type        = string
  default     = ""
}

variable "servicenow_credentials_ref" {
  description = "Secret with ServiceNow credentials, e.g. account.servicenow. The ServiceNow integration is skipped when empty."
  type        = string
  default     = ""
}

# Custom integration at account scope: api mode by default, webhook mode with an HMAC secret.
resource "harness_platform_idp_integration" "orders" {
  identifier  = "${var.prefix}_orders"
  name        = "TF e2e orders"
  description = "Created by examples/idp_integrations"
  type        = "Custom"
  mode        = var.webhook_secret_ref == "" ? "api" : "webhook"
  source      = "orders"
  tags        = ["team:idp", "terraform"]

  dynamic "webhook_auth" {
    for_each = var.webhook_secret_ref == "" ? [] : [var.webhook_secret_ref]
    content {
      auth_type   = "hmac_auth"
      hash_algo   = "hmac_sha256"
      auth_header = "X-Signature"
      secret_ref  = webhook_auth.value
    }
  }

  kind {
    kind        = "custom"
    auto_import = true
    action      = "merge"
  }
}

resource "harness_platform_idp_integration_kind_mapping" "orders" {
  integration_id = harness_platform_idp_integration.orders.identifier
  kind           = "custom"

  ingestion {
    identifier = ".service.id"
    name       = ".service.name"
    properties = {
      "spec.owner" = ".service.team"
    }
  }
}

# Optional Harness Kubernetes services integration; needs a service discovery agent.
resource "harness_platform_idp_integration" "k8s" {
  count = var.k8s_agent_ref == "" ? 0 : 1

  identifier = "${var.prefix}_k8s"
  name       = "TF e2e Harness K8s"
  org_id     = var.org_id
  project_id = var.project_id == "" ? null : var.project_id
  type       = "HarnessK8s"

  harness_k8s {
    service_discovery_agent_identifier = var.k8s_agent_ref
    environment_identifier             = var.k8s_environment
  }

  kind {
    kind = "service"
  }
}

# Optional ServiceNow CMDB integration; needs a real credentials secret.
resource "harness_platform_idp_integration" "servicenow" {
  count = var.servicenow_credentials_ref == "" ? 0 : 1

  identifier       = "${var.prefix}_servicenow"
  name             = "TF e2e ServiceNow"
  org_id           = var.org_id
  type             = "ServiceNow"
  update_frequency = "12h"

  connector {
    inline {
      credentials_ref = var.servicenow_credentials_ref
    }
  }

  kind {
    kind             = "cmdb_ci_service"
    replication_mode = "incremental"
    idp_kind         = "component"
    idp_type         = "service"
  }
}

data "harness_platform_idp_integration_type" "custom" {
  type = "Custom"
}

data "harness_platform_idp_integration" "orders" {
  identifier = harness_platform_idp_integration.orders.identifier
  depends_on = [harness_platform_idp_integration_kind_mapping.orders]
}

data "harness_platform_idp_integration_kind_mapping" "orders" {
  integration_id = harness_platform_idp_integration.orders.identifier
  kind           = "custom"
  depends_on     = [harness_platform_idp_integration_kind_mapping.orders]
}

data "harness_platform_idp_integrations" "org" {
  org_id     = var.org_id
  project_id = var.project_id == "" ? null : var.project_id
  depends_on = [harness_platform_idp_integration.k8s, harness_platform_idp_integration.servicenow]
}

data "harness_platform_idp_integration_catalog" "servicenow" {
  count = var.servicenow_credentials_ref == "" ? 0 : 1

  integration_id = harness_platform_idp_integration.servicenow[0].identifier
  org_id         = var.org_id
}

output "webhook_urls" {
  value = data.harness_platform_idp_integration.orders.endpoints[*].url
}

output "orders_revision" {
  value = data.harness_platform_idp_integration.orders.revision
}

output "orders_mapping_identifier" {
  value = data.harness_platform_idp_integration_kind_mapping.orders.ingestion[0].identifier
}

output "custom_kinds" {
  value = data.harness_platform_idp_integration_type.custom.kinds[*].kind
}

output "org_integrations" {
  value = data.harness_platform_idp_integrations.org.integrations[*].identifier
}

output "servicenow_streams" {
  value = [for c in data.harness_platform_idp_integration_catalog.servicenow : length(c.streams)]
}
