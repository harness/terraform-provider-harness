# ============================================================================
# Harness Chaos Experiment Resource Examples
# ============================================================================
#
# Chaos Experiments are INSTANCES created FROM Experiment Templates.
# These examples are based on TESTED configurations from the e2e-test suite.
#
# Key Points:
# - Experiments are created by "launching" an experiment template
# - Infrastructure binding (infra_ref) is MANDATORY (format: env_id/infra_id)
# - Hub scope can be different from experiment scope (cross-scope support)
# - import_type: "REFERENCE" (default) or "LOCAL" (full copy)
# ============================================================================

# ----------------------------------------------------------------------------
# Example 1: Basic Experiment with REFERENCE Import
# ----------------------------------------------------------------------------
# Most common pattern: project-level experiment from project-level template

resource "harness_chaos_experiment" "basic" {
  depends_on = [
    harness_chaos_experiment_template.my_template,
    harness_chaos_infrastructure_v2.my_infra
  ]

  # Experiment scope
  org_id     = "my_org"
  project_id = "my_project"

  # Hub scope (same as experiment scope in this example)
  hub_org_id     = "my_org"
  hub_project_id = "my_project"
  hub_identity   = "my-chaos-hub"

  # Template and infrastructure
  template_identity = "my-template"
  infra_ref         = "my-env/my-infra"

  # Experiment details
  name        = "My Chaos Experiment"
  identity    = "my-chaos-experiment"
  description = "Basic chaos experiment with REFERENCE import"

  # REFERENCE import: changes to template will reflect in experiment
  import_type = "REFERENCE"

  tags = ["chaos", "kubernetes", "resilience"]

  lifecycle {
    ignore_changes = [tags]
  }
}

# ----------------------------------------------------------------------------
# Example 2: Experiment with LOCAL Import
# ----------------------------------------------------------------------------
# LOCAL import creates a full copy - independent of template changes

resource "harness_chaos_experiment" "local_copy" {
  depends_on = [
    harness_chaos_experiment_template.my_template,
    harness_chaos_infrastructure_v2.my_infra
  ]

  org_id     = "my_org"
  project_id = "my_project"

  hub_org_id     = "my_org"
  hub_project_id = "my_project"
  hub_identity   = "my-chaos-hub"

  template_identity = "my-template"
  infra_ref         = "my-env/my-infra"

  name        = "My Independent Experiment"
  identity    = "my-independent-experiment"
  description = "Experiment with LOCAL import (independent copy)"

  # LOCAL import: full copy, changes to template won't affect this
  import_type = "LOCAL"

  tags = ["chaos", "local", "independent"]

  lifecycle {
    ignore_changes = [tags]
  }
}

# ----------------------------------------------------------------------------
# Example 3: Cross-Scope Experiment (Org Hub → Project Experiment)
# ----------------------------------------------------------------------------
# Create project experiment from org-level template (cross-scope)

resource "harness_chaos_experiment" "cross_scope" {
  depends_on = [
    harness_chaos_experiment_template.org_template,
    harness_chaos_infrastructure_v2.my_infra
  ]

  # Experiment scope (project level)
  org_id     = "my_org"
  project_id = "my_project"

  # Hub scope (org level) - different from experiment scope!
  hub_org_id     = "my_org"
  hub_identity   = "my-chaos-hub"

  # Template from org-level hub
  template_identity = "my-template"
  infra_ref         = "my-env/my-infra"

  name        = "Cross-Scope Experiment"
  identity    = "cross-scope-experiment"
  description = "Project experiment from org-level template"

  import_type = "REFERENCE"

  tags = ["chaos", "cross-scope", "org-template"]

  lifecycle {
    ignore_changes = [tags]
  }
}

# ----------------------------------------------------------------------------
# Key Fields Reference
# ----------------------------------------------------------------------------
# Required:
#   - org_id, project_id (experiment scope)
#   - hub_identity (hub name)
#   - template_identity (template to use)
#   - name (experiment name)
#   - infra_ref (infrastructure binding - format: env_id/infra_id)
#
# Optional:
#   - hub_org_id, hub_project_id (hub scope, defaults to experiment scope)
#   - identity (auto-generated if not provided)
#   - description
#   - import_type ("REFERENCE" or "LOCAL", default: "REFERENCE")
#   - tags
#   - revision (template revision, default: "v1")
#
# Computed (read-only):
#   - experiment_id, infra_id, infra_type
#   - manifest (for LOCAL imports)
#   - template_details (nested block)
#   - created_at, updated_at, created_by, updated_by
#
# Import Format: org_id/project_id/experiment_identity
# Example: terraform import harness_chaos_experiment.basic my_org/my_project/my-experiment
# ----------------------------------------------------------------------------
