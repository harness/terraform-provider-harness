# ============================================================================
# Harness Chaos Action Template Resource Examples
# ============================================================================
#
# Action templates define reusable actions for chaos experiments.
# These examples are based on TESTED configurations from the e2e-test suite.
#
# Key Points:
# - Actions can be delay, script, or container type
# - Runtime inputs with defaults: "<+input>.default('value')"
# - Container actions support full Kubernetes pod configuration
# ============================================================================

# ----------------------------------------------------------------------------
# Example 1: Container Action with Runtime Inputs (TESTED ✅)
# ----------------------------------------------------------------------------
# Most common pattern: container action with runtime inputs and defaults

resource "harness_chaos_action_template" "container_with_runtime_inputs" {
  depends_on = [harness_chaos_hub_v2.project_level]

  # Project level
  org_id       = harness_platform_organization.this.id
  project_id   = harness_platform_project.this.id
  hub_identity = harness_chaos_hub_v2.project_level.identity

  identity            = "container-action-template"
  name                = "Container Action Template"
  description         = "Container action with runtime inputs and defaults"
  type                = "container"
  infrastructure_type = "<+input>.default('Kubernetes')"
  tags                = ["container", "kubernetes", "runtime-inputs"]

  # Container action
  container_action {
    image     = "<+input>.default('busybox:latest')"
    command   = ["<+input>.default('sh')"]
    args      = "echo 'Running container action'; sleep 15"
    namespace = "<+input>.default('default')"

    node_selector = {
      disktype = "ssd"
      zone     = "us-west-1a"
    }

    labels = {
      app         = "chaos-action"
      environment = "production"
      managed-by  = "terraform"
    }

    annotations = {
      description = "Chaos container action"
      owner       = "chaos-team"
    }

    env {
      name  = "TEST_VAR"
      value = "<+input>.default('test_value')"
    }

    env {
      name  = "ANOTHER_VAR"
      value = "<+input>.default('another_value')"
    }

    resources {
      limits = {
        cpu    = "500m"
        memory = "512Mi"
      }

      requests = {
        cpu    = "250m"
        memory = "256Mi"
      }
    }
  }

  # Run properties
  run_properties {
    timeout  = "<+input>.default('60s')"
    interval = "<+input>.default('15s')"
  }

  # Variables
  variables {
    name        = "container_image"
    value       = "<+input>"
    type        = "string"
    required    = true
    description = "Container image to use (runtime input)"
  }

  variables {
    name        = "namespace"
    value       = "<+input>"
    type        = "string"
    required    = false
    description = "Kubernetes namespace (runtime input)"
  }
}

# ----------------------------------------------------------------------------
# Example 2: Simple Delay Action (TESTED ✅)
# ----------------------------------------------------------------------------
# Delay action for adding wait time in experiments

resource "harness_chaos_action_template" "delay_action" {
  depends_on = [harness_chaos_hub_v2.project_level]

  org_id       = harness_platform_organization.this.id
  project_id   = harness_platform_project.this.id
  hub_identity = harness_chaos_hub_v2.project_level.identity

  identity            = "delay-action-template"
  name                = "Delay Action Template"
  description         = "Simple delay action for wait time"
  type                = "delay"
  infrastructure_type = "Kubernetes"
  tags                = ["delay", "wait"]

  # Delay action
  delay_action {
    duration = "<+input>.default('30s')"
  }

  # Run properties
  run_properties {
    timeout = "60s"
  }
}

# ----------------------------------------------------------------------------
# Example 3: Script Action (TESTED ✅)
# ----------------------------------------------------------------------------
# Custom script action for flexible operations

resource "harness_chaos_action_template" "script_action" {
  depends_on = [harness_chaos_hub_v2.project_level]

  org_id       = harness_platform_organization.this.id
  project_id   = harness_platform_project.this.id
  hub_identity = harness_chaos_hub_v2.project_level.identity

  identity            = "script-action-template"
  name                = "Script Action Template"
  description         = "Custom script action for chaos operations"
  type                = "script"
  infrastructure_type = "<+input>.default('Kubernetes')"
  tags                = ["script", "custom"]

  # Custom script action
  custom_script_action {
    script = <<-EOT
      #!/bin/bash
      echo "Running custom chaos script"
      echo "Target: <+input>"
      sleep 10
      echo "Script completed"
    EOT
    
    shell = "bash"
    
    env {
      name  = "TARGET"
      value = "<+input>.default('default-target')"
    }
  }

  # Run properties
  run_properties {
    timeout  = "<+input>.default('120s')"
    interval = "30s"
  }

  # Variables
  variables {
    name        = "target_resource"
    value       = "<+input>"
    type        = "string"
    required    = true
    description = "Target resource for the script"
  }
}

# ----------------------------------------------------------------------------
# Key Fields Reference
# ----------------------------------------------------------------------------
# Required:
#   - org_id, project_id (scope)
#   - hub_identity (hub where template is stored)
#   - identity (template identifier)
#   - name (template display name)
#   - type ("delay", "script", or "container")
#   - infrastructure_type (e.g., "Kubernetes", "KubernetesV2")
#   - One of: delay_action, custom_script_action, or container_action
#
# Optional:
#   - description
#   - tags
#   - run_properties (timeout, interval, polling_interval, etc.)
#   - variables (name, value, type, required, description)
#
# Runtime Inputs:
#   - Use "<+input>" for required runtime input
#   - Use "<+input>.default('value')" for optional with default
#
# Computed (read-only):
#   - action_id, created_at, updated_at
#
# Import Format: org_id/project_id/hub_identity/identity
# Example: terraform import harness_chaos_action_template.example \
#          my_org/my_project/my-hub/my-action
# ----------------------------------------------------------------------------
