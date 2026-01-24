# ============================================================================
# Harness Chaos Experiment Template Resource Examples
# ============================================================================
#
# Experiment templates define complete chaos experiments with actions, faults, and probes.
# These examples are based on TESTED configurations from the e2e-test suite.
#
# Key Points:
# - Templates combine actions, faults, and probes into workflows
# - Vertices define execution order (workflow graph)
# - Runtime inputs supported: "<+input>" or "<+input>.default('value')"
# - Enterprise features: is_enterprise = true
# ============================================================================

# ----------------------------------------------------------------------------
# Example 1: Simple Fault Template (TESTED ✅)
# ----------------------------------------------------------------------------
# Basic template with single fault

resource "harness_chaos_experiment_template" "simple_fault" {
  depends_on = [harness_chaos_hub_v2.project_level]

  org_id       = harness_platform_organization.this.id
  project_id   = harness_platform_project.this.id
  hub_identity = harness_chaos_hub_v2.project_level.identity

  identity    = "simple-pod-delete"
  name        = "Simple Pod Delete Experiment"
  description = "Basic experiment with single pod delete fault"
  tags        = ["kubernetes", "pod-delete", "simple"]

  spec {
    infra_type = "KubernetesV2"

    faults {
      identity      = "pod-delete"
      name          = "pod-delete-fault"
      revision      = "v1"
      is_enterprise = true
      auth_enabled  = false

      values {
        name  = "TARGET_WORKLOAD_KIND"
        value = "deployment"
      }

      values {
        name  = "TARGET_WORKLOAD_NAMESPACE"
        value = "<+input>.default('default')"
      }

      values {
        name  = "TOTAL_CHAOS_DURATION"
        value = "<+input>.default('30s')"
      }
    }

    vertices {
      name = "pod-delete-vertex"
      start {
        faults {
          name = "pod-delete-fault"
        }
      }
      end {}
    }

    cleanup_policy = "delete"
  }
}

# ----------------------------------------------------------------------------
# Example 2: Template with Action and Fault (TESTED ✅)
# ----------------------------------------------------------------------------
# Template combining action and fault

resource "harness_chaos_experiment_template" "with_action" {
  depends_on = [harness_chaos_hub_v2.project_level]

  org_id       = harness_platform_organization.this.id
  project_id   = harness_platform_project.this.id
  hub_identity = harness_chaos_hub_v2.project_level.identity

  identity    = "action-and-fault"
  name        = "Action and Fault Experiment"
  description = "Experiment with action before fault"
  tags        = ["kubernetes", "action", "fault"]

  spec {
    infra_type = "KubernetesV2"

    # Action: Pre-chaos notification
    actions {
      identity               = "notification-action"
      name                   = "pre-chaos-notification"
      is_enterprise          = false
      continue_on_completion = false

      values {
        name  = "MESSAGE"
        value = "Starting chaos experiment"
      }
    }

    # Fault: Container kill
    faults {
      identity      = "container-kill"
      name          = "container-kill-fault"
      revision      = "v1"
      is_enterprise = true
      auth_enabled  = false

      values {
        name  = "TARGET_WORKLOAD_KIND"
        value = "deployment"
      }

      values {
        name  = "TARGET_WORKLOAD_NAMESPACE"
        value = "<+input>"
      }

      values {
        name  = "TOTAL_CHAOS_DURATION"
        value = "<+input>.default('30s')"
      }
    }

    # Workflow: Action first, then fault
    vertices {
      name = "action-vertex"
      start {
        actions {
          name = "pre-chaos-notification"
        }
      }
      end {}
    }

    vertices {
      name = "fault-vertex"
      start {
        faults {
          name = "container-kill-fault"
        }
      }
      end {}
    }

    cleanup_policy = "delete"
  }
}

# ----------------------------------------------------------------------------
# Example 3: Complex Template with Probes (TESTED ✅)
# ----------------------------------------------------------------------------
# Complete template with actions, faults, and probes

resource "harness_chaos_experiment_template" "complex" {
  depends_on = [harness_chaos_hub_v2.project_level]

  org_id       = harness_platform_organization.this.id
  project_id   = harness_platform_project.this.id
  hub_identity = harness_chaos_hub_v2.project_level.identity

  identity    = "complex-experiment"
  name        = "Complex Chaos Experiment"
  description = "Complete experiment with actions, faults, and probes"
  tags        = ["kubernetes", "complex", "enterprise"]

  spec {
    infra_type = "KubernetesV2"

    # Action: Start notification
    actions {
      identity               = "notification-action"
      name                   = "start-notification"
      is_enterprise          = false
      continue_on_completion = false

      values {
        name  = "MESSAGE"
        value = "Chaos experiment started"
      }
    }

    # Fault 1: Pod delete
    faults {
      identity      = "pod-delete"
      name          = "pod-delete-fault"
      revision      = "v1"
      is_enterprise = true
      auth_enabled  = false

      values {
        name  = "TARGET_WORKLOAD_KIND"
        value = "deployment"
      }

      values {
        name  = "TARGET_WORKLOAD_NAMESPACE"
        value = "<+input>"
      }

      values {
        name  = "TOTAL_CHAOS_DURATION"
        value = "<+input>.default('30s')"
      }
    }

    # Fault 2: Network latency
    faults {
      identity      = "pod-network-latency"
      name          = "network-latency-fault"
      revision      = "v1"
      is_enterprise = true
      auth_enabled  = false

      values {
        name  = "TARGET_WORKLOAD_KIND"
        value = "deployment"
      }

      values {
        name  = "TARGET_WORKLOAD_NAMESPACE"
        value = "<+input>"
      }

      values {
        name  = "NETWORK_LATENCY"
        value = "<+input>.default('2000')"
      }
    }

    # Probe 1: Pod status check
    probes {
      identity               = "pod-status-check"
      name                   = "pod-status-probe"
      revision               = "v1"
      is_enterprise          = true
      duration               = 30
      weightage              = 10
      enable_data_collection = false

      conditions = ["onChaosStart", "duringChaos", "afterChaos"]

      values {
        name  = "TARGET_NAMESPACE"
        value = "<+input>"
      }
    }

    # Probe 2: HTTP health check
    probes {
      identity               = "http-health-check"
      name                   = "http-health-probe"
      revision               = "v1"
      is_enterprise          = true
      duration               = 30
      weightage              = 10
      enable_data_collection = false

      conditions = ["duringChaos", "afterChaos"]

      values {
        name  = "URL"
        value = "<+input>"
      }
    }

    # Workflow: 3 stages
    vertices {
      name = "action-stage"
      start {
        actions {
          name = "start-notification"
        }
      }
      end {}
    }

    vertices {
      name = "fault-stage"
      start {
        faults {
          name = "pod-delete-fault"
        }
        faults {
          name = "network-latency-fault"
        }
        probes {
          name = "pod-status-probe"
        }
        probes {
          name = "http-health-probe"
        }
      }
      end {}
    }

    vertices {
      name = "cleanup-stage"
      start {}
      end {}
    }

    cleanup_policy = "delete"

    status_check_timeouts {
      delay   = 5
      timeout = 300
    }
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
#   - spec.infra_type (e.g., "KubernetesV2")
#   - spec.vertices (workflow graph)
#
# Optional:
#   - description
#   - tags
#   - spec.actions (action templates to use)
#   - spec.faults (fault templates to use)
#   - spec.probes (probe templates to use)
#   - spec.cleanup_policy ("delete" or "retain")
#   - spec.status_check_timeouts (delay, timeout)
#
# Actions/Faults/Probes:
#   - identity: Template identity to reference
#   - name: Instance name in workflow
#   - revision: Template version (default: "v1")
#   - is_enterprise: Use enterprise features
#   - values: Parameter overrides (name, value)
#
# Probes Additional Fields:
#   - duration: Probe duration in seconds
#   - weightage: Probe importance (0-100)
#   - enable_data_collection: Collect probe data
#   - conditions: When to run (onChaosStart, duringChaos, afterChaos)
#
# Vertices (Workflow):
#   - name: Stage name
#   - start: Resources to execute at start (actions, faults, probes)
#   - end: Resources to execute at end (usually empty)
#
# Runtime Inputs:
#   - Use "<+input>" for required runtime input
#   - Use "<+input>.default('value')" for optional with default
#
# Computed (read-only):
#   - template_id, created_at, updated_at
#
# Import Format: org_id/project_id/hub_identity/identity
# Example: terraform import harness_chaos_experiment_template.example \
#          my_org/my_project/my-hub/my-template
# ----------------------------------------------------------------------------
