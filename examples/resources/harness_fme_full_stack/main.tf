# End-to-end exercise of Harness FME (Split) data sources and resources.
#
# Prerequisites:
#   - HARNESS_ACCOUNT_ID, HARNESS_PLATFORM_API_KEY (Split admin API), and platform/org permissions to create orgs/projects.
#   - Provider version that includes the harness_fme_* data sources and resources.
#
# Usage:
#   terraform init
#   terraform apply -var="name_prefix=mydemo"
#
# Optional: set -var="skip_api_key=true" to skip harness_fme_api_key (avoids minting a real Split API key).

terraform {
  required_version = ">= 1.3.0"
  required_providers {
    harness = {
      source  = "harness/harness"
      version = ">= 0.35.0"
    }
    random = {
      source  = "hashicorp/random"
      version = ">= 3.5.0"
    }
  }
}

variable "name_prefix" {
  type        = string
  description = "Short alphanumeric prefix for generated Harness org/project identifiers."
  default     = "fmestack"
}

variable "skip_api_key" {
  type        = bool
  description = "If true, do not create harness_fme_api_key (avoids minting a real Split API key)."
  default     = false
}

provider "harness" {}

resource "random_string" "suffix" {
  length  = 6
  lower   = true
  upper   = false
  special = false
}

locals {
  harness_id = "${var.name_prefix}_${random_string.suffix.result}"
  # Keys for each Split environment we create (names derived below, max 20 chars for Split).
  fme_environment_keys = toset(["alpha", "beta"])
  fme_environment_names = {
    for k in local.fme_environment_keys : k => substr("tf${k}_${random_string.suffix.result}", 0, 20)
  }

  segment_name  = "tfseg_${random_string.suffix.result}"
  rbs_name      = "tfrbs_${random_string.suffix.result}"
  ls_name       = "tf_ls_${random_string.suffix.result}"
  flag_name     = "tf_flag_${random_string.suffix.result}"
  flag_set_name = "tf_fs_${random_string.suffix.result}"
  attr_id       = "tf_attr_${random_string.suffix.result}"
  tt_name       = "tftt_${random_string.suffix.result}"

  # Split SplitDefinitionRequest shape: multi-treatment default rollout + a targeted rule with buckets.
  # See https://docs.split.io/reference/create-feature-flag-definition-in-environment
  feature_definition = jsonencode({
    title   = "Full-stack example (complex definition)"
    comment = "Multi-treatment defaultRule + whitelist-style targeting rule (Terraform demo)."
    treatments = [
      {
        name           = "on"
        description    = "Enabled"
        configurations = jsonencode({ variant = "full", schemaVersion = 1 })
      },
      { name = "beta", description = "Partial / beta cohort" },
      { name = "off", description = "Disabled" },
    ]
    baselineTreatment = "off"
    defaultTreatment  = "off"
    trafficAllocation = 100
    defaultRule = [
      { treatment = "on", size = 15 },
      { treatment = "beta", size = 35 },
      { treatment = "off", size = 50 },
    ]
    rules = [
      {
        condition = {
          combiner = "AND"
          matchers = [
            {
              type    = "IN_LIST_STRING"
              strings = ["demo_vip_user", "demo_partner_account"]
            }
          ]
        }
        buckets = [
          { treatment = "on", size = 100 }
        ]
      }
    ]
  })

  # Env-scoped RBS definition shape matches Split list/get: only excludedKeys + rules[].condition.matchers[].
  # Do not send response-only fields (id, environment, trafficType, changeNumber). title/comment are often
  # absent on env payloads; omit them here to match working APIs. Use IN_LIST_STRING only unless the traffic
  # type actually defines attributes for CONTAINS_STRING / etc. (your sample uses traffic type "account" + keys).
  rbs_definition = jsonencode({
    excludedKeys = ["excluded_qa_bot", "excluded_loadgen"]
    rules = [
      {
        condition = {
          combiner = "AND"
          matchers = [
            {
              type    = "IN_LIST_STRING"
              strings = ["cohort_demo_1", "cohort_demo_2"]
            }
          ]
        }
      }
    ]
  })
}

resource "harness_platform_organization" "this" {
  identifier = local.harness_id
  name       = local.harness_id
}

resource "harness_platform_project" "this" {
  identifier = local.harness_id
  org_id     = harness_platform_organization.this.id
  name       = local.harness_id
}

# --- Data sources: workspace, default workspace traffic type, flag set (after create) ---

data "harness_fme_workspace" "by_project" {
  org_id     = harness_platform_organization.this.id
  project_id = harness_platform_project.this.id
}

data "harness_fme_workspace" "by_name" {
  name = data.harness_fme_workspace.by_project.name
}

data "harness_fme_traffic_type" "user" {
  org_id     = harness_platform_organization.this.id
  project_id = harness_platform_project.this.id
  name       = "user"
}

resource "harness_fme_flag_set" "this" {
  org_id      = harness_platform_organization.this.id
  project_id  = harness_platform_project.this.id
  name        = local.flag_set_name
  description = "Terraform FME full-stack example"
}

data "harness_fme_flag_set" "lookup" {
  depends_on = [harness_fme_flag_set.this]
  org_id     = harness_platform_organization.this.id
  project_id = harness_platform_project.this.id
  name       = harness_fme_flag_set.this.name
}

# --- Resources ---

resource "harness_fme_environment" "stack" {
  for_each   = local.fme_environment_keys
  org_id     = harness_platform_organization.this.id
  project_id = harness_platform_project.this.id
  name       = local.fme_environment_names[each.key]
  production = false
  depends_on = [harness_platform_project.this]
}

resource "harness_fme_traffic_type" "extra" {
  org_id     = harness_platform_organization.this.id
  project_id = harness_platform_project.this.id
  name       = local.tt_name
  depends_on = [harness_platform_project.this]
}

resource "harness_fme_traffic_type_attribute" "example" {
  org_id          = harness_platform_organization.this.id
  project_id      = harness_platform_project.this.id
  traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
  identifier      = local.attr_id
  display_name    = "TF Example Attribute"
  data_type       = "string"
  is_searchable   = false
}

resource "harness_fme_segment" "classic" {
  org_id          = harness_platform_organization.this.id
  project_id      = harness_platform_project.this.id
  traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
  name            = local.segment_name
  description     = "classic segment from Terraform example"
}

resource "harness_fme_segment_environment_association" "classic" {
  for_each       = local.fme_environment_keys
  org_id         = harness_platform_organization.this.id
  project_id     = harness_platform_project.this.id
  environment_id = harness_fme_environment.stack[each.key].environment_id
  segment_name   = harness_fme_segment.classic.name
  depends_on     = [harness_fme_segment.classic, harness_fme_environment.stack]
}

resource "harness_fme_environment_segment_keys" "classic" {
  for_each       = local.fme_environment_keys
  org_id         = harness_platform_organization.this.id
  project_id     = harness_platform_project.this.id
  environment_id = harness_fme_environment.stack[each.key].environment_id
  segment_name   = harness_fme_segment.classic.name
  keys           = ["example_user_1", "example_user_2"]
  depends_on     = [harness_fme_segment_environment_association.classic]
}

resource "harness_fme_feature_flag" "toggle" {
  org_id          = harness_platform_organization.this.id
  project_id      = harness_platform_project.this.id
  traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
  name            = local.flag_name
  description     = "Terraform FME example flag"
  tags            = ["terraform-example", "fme-stack"]
}

resource "harness_fme_feature_flag_definition" "toggle" {
  for_each       = local.fme_environment_keys
  org_id         = harness_platform_organization.this.id
  project_id     = harness_platform_project.this.id
  environment_id = harness_fme_environment.stack[each.key].environment_id
  flag_name      = harness_fme_feature_flag.toggle.name
  definition     = local.feature_definition
  depends_on     = [harness_fme_feature_flag.toggle, harness_fme_environment.stack]
}

resource "harness_fme_rule_based_segment" "example" {
  org_id          = harness_platform_organization.this.id
  project_id      = harness_platform_project.this.id
  traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
  name            = local.rbs_name
  depends_on      = [harness_platform_project.this]
}

resource "harness_fme_rule_based_segment_environment_association" "example" {
  for_each        = local.fme_environment_keys
  org_id          = harness_platform_organization.this.id
  project_id      = harness_platform_project.this.id
  environment_id  = harness_fme_environment.stack[each.key].environment_id
  segment_name    = harness_fme_rule_based_segment.example.name
  definition_json = local.rbs_definition
  depends_on      = [harness_fme_rule_based_segment.example, harness_fme_environment.stack]
}

resource "harness_fme_large_segment" "example" {
  org_id          = harness_platform_organization.this.id
  project_id      = harness_platform_project.this.id
  traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
  name            = local.ls_name
  description     = "large segment from Terraform example"
  depends_on      = [harness_platform_project.this]
}

resource "harness_fme_large_segment_environment_association" "example" {
  for_each       = local.fme_environment_keys
  org_id         = harness_platform_organization.this.id
  project_id     = harness_platform_project.this.id
  environment_id = harness_fme_environment.stack[each.key].environment_id
  segment_name   = harness_fme_large_segment.example.name
  depends_on     = [harness_fme_large_segment.example, harness_fme_environment.stack]
}

resource "harness_fme_api_key" "optional" {
  for_each = var.skip_api_key ? toset([]) : local.fme_environment_keys

  org_id         = harness_platform_organization.this.id
  project_id     = harness_platform_project.this.id
  name           = "tf_${each.key}_${random_string.suffix.result}"
  api_key_type   = "server_side"
  environment_id = harness_fme_environment.stack[each.key].environment_id
  depends_on     = [harness_fme_environment.stack]
}

output "fme_workspace_id" {
  value = data.harness_fme_workspace.by_project.workspace_id
}

output "fme_flag_set_id_from_data" {
  value = data.harness_fme_flag_set.lookup.flag_set_id
}

# Map of logical key -> Split environment id for each `harness_fme_environment.stack` instance.
output "fme_environment_ids" {
  value = { for k, env in harness_fme_environment.stack : k => env.environment_id }
}

output "fme_api_key_secrets" {
  value     = { for k, key in harness_fme_api_key.optional : k => key.api_key }
  sensitive = true
}
