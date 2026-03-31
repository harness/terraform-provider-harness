# ============================================================================
# Harness Chaos Experiment Data Source Examples
# ============================================================================
#
# The harness_chaos_experiment data source retrieves information about
# existing chaos experiments.
#
# Lookup Methods:
# 1. By identity (fast, direct GET) - RECOMMENDED
# 2. By name (slower, uses LIST + filter)
# ============================================================================

# ----------------------------------------------------------------------------
# Example 1: Lookup by Identity (Recommended)
# ----------------------------------------------------------------------------
# Fast lookup using the experiment's unique identity

data "harness_chaos_experiment" "by_identity" {
  org_id     = "my_org"
  project_id = "my_project"
  identity   = "my-chaos-experiment"
}

# Use the retrieved data
output "experiment_id" {
  value       = data.harness_chaos_experiment.by_identity.experiment_id
  description = "Internal experiment ID"
}

output "experiment_infra_type" {
  value       = data.harness_chaos_experiment.by_identity.infra_type
  description = "Infrastructure type (e.g., KubernetesV2)"
}

output "experiment_template_name" {
  value       = data.harness_chaos_experiment.by_identity.template_details[0].template_name
  description = "Name of the template used"
}

# ----------------------------------------------------------------------------
# Example 2: Lookup by Name
# ----------------------------------------------------------------------------
# Lookup using the experiment's display name (slower than identity lookup)

data "harness_chaos_experiment" "by_name" {
  org_id     = "my_org"
  project_id = "my_project"
  name       = "My Chaos Experiment"
}

output "experiment_identity" {
  value       = data.harness_chaos_experiment.by_name.identity
  description = "Experiment identity"
}

# ----------------------------------------------------------------------------
# Example 3: Use in Other Resources
# ----------------------------------------------------------------------------
# Reference experiment data in other resources

data "harness_chaos_experiment" "existing" {
  org_id     = "my_org"
  project_id = "my_project"
  identity   = "prod-experiment"
}

# Example: Create monitoring based on experiment
resource "harness_platform_monitored_service" "chaos_monitor" {
  org_identifier     = data.harness_chaos_experiment.existing.org_id
  project_identifier = data.harness_chaos_experiment.existing.project_id
  identifier         = "monitor-${data.harness_chaos_experiment.existing.identity}"
  name               = "Monitor for ${data.harness_chaos_experiment.existing.name}"
  description        = "Monitoring for: ${data.harness_chaos_experiment.existing.description}"
}

# ----------------------------------------------------------------------------
# Available Attributes
# ----------------------------------------------------------------------------
# Inputs (required):
#   - org_id, project_id
#   - identity OR name (one required)
#
# Outputs (computed):
#   - id, experiment_id, identity, name, description
#   - infra_ref, infra_id, infra_type
#   - manifest (for LOCAL imports)
#   - template_details (nested block with template info)
#   - tags, created_at, updated_at
#   - created_by, updated_by, is_removed
#
# Best Practices:
#   - Prefer identity lookup (faster)
#   - Use for read-only access to experiments
#   - Mark sensitive data (manifest) as sensitive in outputs
# ----------------------------------------------------------------------------
