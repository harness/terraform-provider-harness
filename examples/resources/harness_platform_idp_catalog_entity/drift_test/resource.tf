terraform {
  required_providers {
    harness = {
      source = "registry.terraform.io/harness/harness"
    }
  }
}

provider "harness" {}

variable "catalog_entity_identifier" {
  type        = string
  description = "Unique IDP catalog entity identifier to create for drift testing."
  default     = "terraform_drift_test_component"
}

resource "harness_platform_idp_catalog_entity" "drift_test" {
  identifier = var.catalog_entity_identifier
  kind       = "component"
  yaml       = <<-EOT
    apiVersion: harness.io/v1
    kind: Component
    type: service
    identifier: ${var.catalog_entity_identifier}
    name: Terraform Drift Test Component
    owner: user:account/admin@harness.io
    spec:
      lifecycle: experimental
    metadata:
      description: Managed by Terraform. Change this description in Harness to test drift detection. 
      tags:
        - terraform
        - drift-test
        - new-tag
    EOT
}
