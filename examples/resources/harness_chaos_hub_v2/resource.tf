# ============================================================================
# Harness Chaos Hub V2 Resource Examples
# ============================================================================
#
# Chaos Hubs store chaos artifacts (templates, faults, probes, actions).
# These examples are based on TESTED configurations from the e2e-test suite.
#
# Key Points:
# - Hubs can be created at account, org, or project scope
# - connector_ref, repo_branch, repo_name are OPTIONAL (not required)
# - Use lifecycle { ignore_changes = [tags] } to prevent drift
# ============================================================================

# ----------------------------------------------------------------------------
# Example 1: Account-Level Chaos Hub
# ----------------------------------------------------------------------------
# Account-level hub accessible across all orgs and projects

resource "harness_chaos_hub_v2" "account_level" {
  # Account level - no org_id or project_id
  identity    = "account-chaos-hub"
  name        = "Account Chaos Hub"
  description = "Account-level chaos hub for enterprise-wide templates"

  tags = ["account", "enterprise", "chaos"]

  lifecycle {
    ignore_changes = [tags]
  }
}

# ----------------------------------------------------------------------------
# Example 2: Org-Level Chaos Hub (TESTED ✅)
# ----------------------------------------------------------------------------
# Org-level hub accessible within the organization

resource "harness_chaos_hub_v2" "org_level" {
  depends_on = [
    harness_platform_organization.this
  ]

  # Org level - org_id only
  org_id      = harness_platform_organization.this.id
  identity    = "org-chaos-hub"
  name        = "Org Chaos Hub"
  description = "Org-level chaos hub for shared templates"

  tags = ["org", "shared", "chaos"]

  lifecycle {
    ignore_changes = [tags]
  }
}

# ----------------------------------------------------------------------------
# Example 3: Project-Level Chaos Hub
# ----------------------------------------------------------------------------
# Project-level hub for project-specific templates

resource "harness_chaos_hub_v2" "project_level" {
  depends_on = [
    harness_platform_project.this
  ]

  # Project level - org_id and project_id
  org_id      = harness_platform_organization.this.id
  project_id  = harness_platform_project.this.id
  identity    = "project-chaos-hub"
  name        = "Project Chaos Hub"
  description = "Project-level chaos hub for team templates"

  tags = ["project", "team", "chaos"]

  lifecycle {
    ignore_changes = [tags]
  }
}

# ----------------------------------------------------------------------------
# Key Fields Reference
# ----------------------------------------------------------------------------
# Required:
#   - identity (hub identifier)
#   - name (hub display name)
#
# Optional:
#   - org_id (for org/project scope)
#   - project_id (for project scope)
#   - description
#   - tags
#   - connector_ref (Git connector for custom hubs)
#   - repo_branch (Git branch)
#   - repo_name (Git repository name)
#
# Computed (read-only):
#   - hub_id (internal hub ID)
#   - created_at, updated_at
#   - is_default
#
# Import Format: org_id/project_id/identity
# Example: terraform import harness_chaos_hub_v2.project my_org/my_project/my-hub
#
# Lifecycle:
#   - Use ignore_changes = [tags] to prevent tag drift
#   - Use depends_on for proper resource ordering
# ----------------------------------------------------------------------------
