# Step 1 — Apply: Run `terraform apply` to create the catalog entity.
# Step 2 — Modify: In the Harness UI, change `metadata.description`, for example to
# "Changed manually in Harness".
# Step 3 — Detect: Run `terraform plan` and verify Terraform reports drift for `yaml`.
# Step 4 — Remediate: Run `terraform apply` to restore the configuration in this file.

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
