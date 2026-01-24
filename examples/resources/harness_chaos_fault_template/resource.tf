# ============================================================================
# Harness Chaos Fault Template Resource Examples
# ============================================================================
#
# Fault templates define reusable chaos faults for experiments.
# These examples are based on TESTED configurations from the e2e-test suite.
#
# Key Points:
# - Faults inject failures into systems (pod delete, network latency, etc.)
# - Type is usually "Custom" for custom faults
# - Category and infrastructures define where fault can run
# ============================================================================

# ----------------------------------------------------------------------------
# Example 1: Basic Kubernetes Fault (TESTED ✅)
# ----------------------------------------------------------------------------
# Most common pattern: Custom Kubernetes fault with container spec

resource "harness_chaos_fault_template" "kubernetes_fault" {
  depends_on = [harness_chaos_hub_v2.project_level]

  # Project level
  org_id       = harness_platform_organization.this.id
  project_id   = harness_platform_project.this.id
  hub_identity = harness_chaos_hub_v2.project_level.identity

  identity             = "k8s-fault-template"
  name                 = "Kubernetes Fault Template"
  description          = "Custom Kubernetes fault for chaos injection"
  category             = ["Kubernetes"]
  infrastructures      = ["KubernetesV2"]
  type                 = "Custom"
  permissions_required = "Basic"
  tags                 = ["kubernetes", "fault", "custom"]

  links {
    name = "Documentation"
    url  = "https://docs.harness.io/chaos"
  }

  spec {
    chaos {
      fault_name = "byoc-injector"

      params {
        name  = "CHAOS_DURATION"
        value = "30s"
      }
      
      params {
        name  = "CHAOS_INTERVAL"
        value = "5s"
      }

      kubernetes {
        image             = "chaosnative/go-runner:ci"
        command           = ["/bin/bash", "-c"]
        args              = ["echo 'Running chaos fault'; sleep 30"]
        image_pull_policy = "IfNotPresent"

        resources {
          limits = {
            cpu    = "150m"
            memory = "150Mi"
          }
          
          requests = {
            cpu    = "100m"
            memory = "100Mi"
          }
        }
      }
    }
  }
}

# ----------------------------------------------------------------------------
# Example 2: Fault with Environment Variables (TESTED ✅)
# ----------------------------------------------------------------------------
# Fault with environment variables for configuration

resource "harness_chaos_fault_template" "fault_with_env" {
  depends_on = [harness_chaos_hub_v2.project_level]

  org_id       = harness_platform_organization.this.id
  project_id   = harness_platform_project.this.id
  hub_identity = harness_chaos_hub_v2.project_level.identity

  identity             = "fault-with-env-template"
  name                 = "Fault with Environment Variables"
  description          = "Fault template with environment configuration"
  category             = ["Kubernetes"]
  infrastructures      = ["KubernetesV2"]
  type                 = "Custom"
  permissions_required = "Basic"
  tags                 = ["kubernetes", "env", "config"]

  links {
    name = "Documentation"
    url  = "https://docs.harness.io/chaos"
  }

  spec {
    chaos {
      fault_name = "byoc-injector"

      params {
        name  = "CHAOS_DURATION"
        value = "15s"
      }
      
      params {
        name  = "CHAOS_INTERVAL"
        value = "3s"
      }
      
      params {
        name  = "TARGET_NAMESPACE"
        value = "<+input>.default('default')"
      }

      kubernetes {
        image             = "chaosnative/go-runner:ci"
        command           = ["/bin/bash", "-c"]
        args              = ["echo 'Fault with env vars'; sleep 15"]
        image_pull_policy = "IfNotPresent"

        env {
          name  = "TARGET_NAMESPACE"
          value = "<+input>.default('default')"
        }
        
        env {
          name  = "CHAOS_MODE"
          value = "pod"
        }

        resources {
          limits = {
            cpu    = "200m"
            memory = "200Mi"
          }
        }
      }
    }
  }

  # Variables
  variables {
    name        = "target_namespace"
    value       = "<+input>"
    type        = "string"
    required    = false
    description = "Target namespace for chaos injection"
  }
}

# ----------------------------------------------------------------------------
# Example 3: Fault with Advanced Configuration (TESTED ✅)
# ----------------------------------------------------------------------------
# Fault with node selector, labels, and annotations

resource "harness_chaos_fault_template" "advanced_fault" {
  depends_on = [harness_chaos_hub_v2.project_level]

  org_id       = harness_platform_organization.this.id
  project_id   = harness_platform_project.this.id
  hub_identity = harness_chaos_hub_v2.project_level.identity

  identity             = "advanced-fault-template"
  name                 = "Advanced Fault Template"
  description          = "Fault with advanced Kubernetes configuration"
  category             = ["Kubernetes"]
  infrastructures      = ["KubernetesV2"]
  type                 = "Custom"
  permissions_required = "Basic"
  tags                 = ["kubernetes", "advanced", "production"]

  links {
    name = "Documentation"
    url  = "https://docs.harness.io/chaos"
  }
  
  links {
    name = "Support"
    url  = "https://support.harness.io"
  }

  spec {
    chaos {
      fault_name = "byoc-injector"

      params {
        name  = "CHAOS_DURATION"
        value = "<+input>.default('30s')"
      }
      
      params {
        name  = "CHAOS_INTERVAL"
        value = "<+input>.default('5s')"
      }

      kubernetes {
        image             = "chaosnative/go-runner:ci"
        command           = ["/bin/bash", "-c"]
        args              = ["echo 'Advanced chaos fault'; sleep 30"]
        image_pull_policy = "IfNotPresent"
        
        node_selector = {
          disktype = "ssd"
          zone     = "us-west-1a"
        }
        
        labels = {
          app         = "chaos-fault"
          environment = "production"
          managed-by  = "terraform"
        }
        
        annotations = {
          description = "Advanced chaos fault"
          owner       = "chaos-team"
        }

        resources {
          limits = {
            cpu    = "250m"
            memory = "256Mi"
          }
          
          requests = {
            cpu    = "125m"
            memory = "128Mi"
          }
        }
      }
    }
  }

  # Variables
  variables {
    name        = "chaos_duration"
    value       = "<+input>"
    type        = "string"
    required    = true
    description = "Duration of chaos injection"
  }
  
  variables {
    name        = "chaos_interval"
    value       = "<+input>"
    type        = "string"
    required    = false
    description = "Interval between chaos injections"
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
#   - category (e.g., ["Kubernetes"])
#   - infrastructures (e.g., ["KubernetesV2"])
#   - type (usually "Custom")
#   - permissions_required ("Basic", "Advanced", "Expert")
#   - spec.chaos (fault configuration)
#
# Optional:
#   - description
#   - tags
#   - links (name, url)
#   - variables (name, value, type, required, description)
#
# Spec.Chaos Fields:
#   - fault_name: Name of the fault (e.g., "byoc-injector")
#   - params: Fault parameters (name, value)
#   - kubernetes: Kubernetes pod specification
#     - image, command, args
#     - image_pull_policy
#     - resources (limits, requests)
#     - env (environment variables)
#     - node_selector, labels, annotations
#
# Runtime Inputs:
#   - Use "<+input>" for required runtime input
#   - Use "<+input>.default('value')" for optional with default
#
# Computed (read-only):
#   - fault_id, created_at, updated_at
#
# Import Format: org_id/project_id/hub_identity/identity
# Example: terraform import harness_chaos_fault_template.example \
#          my_org/my_project/my-hub/my-fault
# ----------------------------------------------------------------------------
