terraform {
  required_providers {
    harness = {
      source = "harness/harness"
    }
  }
}

resource "harness_platform_resource_group" "resource_group" {
    identifier = replace(lower(var.name), " ", "_")
    name = var.name
    description = "resource group to manage ${var.name}"
    tags = var.tags

    account_id = account_id
    org_id = organization_id
    allowed_scope_levels = ["organization"]
    included_scopes {
        filter= [INCLUDING_CHILD_SCOPES]
        account_id = account_id
        org_id = organization_id
    }

    resource_filter {
        include_all_resources = true
    }
}