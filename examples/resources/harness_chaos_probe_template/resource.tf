# ============================================================================
# Harness Chaos Probe Template Resource Examples
# ============================================================================
#
# Probe templates define health checks for chaos experiments.
# These examples are based on TESTED configurations from the e2e-test suite.
#
# Key Points:
# - Probe types: httpProbe, cmdProbe, k8sProbe, promProbe
# - Runtime inputs supported: "<+input>.default('value')"
# - Run properties control probe behavior (timeout, interval, etc.)
# ============================================================================

# ----------------------------------------------------------------------------
# Example 1: K8s Probe (TESTED ✅)
# ----------------------------------------------------------------------------
# Most common pattern: Kubernetes resource probe

resource "harness_chaos_probe_template" "k8s_probe" {
  depends_on = [harness_chaos_hub_v2.project_level]

  # Project level
  org_id       = harness_platform_organization.this.id
  project_id   = harness_platform_project.this.id
  hub_identity = harness_chaos_hub_v2.project_level.identity

  identity            = "k8s-probe-template"
  name                = "K8s Probe Template"
  description         = "Kubernetes resource probe for deployment health"
  type                = "k8sProbe"
  infrastructure_type = "KubernetesV2"
  tags                = ["kubernetes", "probe", "health-check"]

  # K8s probe configuration
  k8s_probe {
    resource  = "deployments"
    namespace = "<+input>.default('default')"
    operation = "present"
    version   = "v1"
  }

  # Run properties
  run_properties {
    timeout         = "15s"
    interval        = "5s"
    stop_on_failure = false
    verbosity       = "INFO"
  }

  # Variables
  variables {
    name        = "target_namespace"
    value       = "<+input>"
    type        = "string"
    required    = false
    description = "Kubernetes namespace to probe"
  }
}

# ----------------------------------------------------------------------------
# Example 2: HTTP Probe (TESTED ✅)
# ----------------------------------------------------------------------------
# HTTP endpoint health check probe

resource "harness_chaos_probe_template" "http_probe" {
  depends_on = [harness_chaos_hub_v2.project_level]

  org_id       = harness_platform_organization.this.id
  project_id   = harness_platform_project.this.id
  hub_identity = harness_chaos_hub_v2.project_level.identity

  identity            = "http-probe-template"
  name                = "HTTP Probe Template"
  description         = "HTTP endpoint health check"
  type                = "httpProbe"
  infrastructure_type = "KubernetesV2"
  tags                = ["http", "probe", "endpoint"]

  # HTTP probe configuration
  http_probe {
    url    = "<+input>.default('http://localhost:8080/health')"
    method = "GET"
    
    headers {
      key   = "Content-Type"
      value = "application/json"
    }
    
    headers {
      key   = "Authorization"
      value = "<+input>"
    }
  }

  # Run properties
  run_properties {
    timeout         = "30s"
    interval        = "10s"
    polling_interval = "2s"
    stop_on_failure = true
    verbosity       = "INFO"
  }

  # Variables
  variables {
    name        = "endpoint_url"
    value       = "<+input>"
    type        = "string"
    required    = true
    description = "HTTP endpoint URL to probe"
  }
}

# ----------------------------------------------------------------------------
# Example 3: CMD Probe (TESTED ✅)
# ----------------------------------------------------------------------------
# Command execution probe

resource "harness_chaos_probe_template" "cmd_probe" {
  depends_on = [harness_chaos_hub_v2.project_level]

  org_id       = harness_platform_organization.this.id
  project_id   = harness_platform_project.this.id
  hub_identity = harness_chaos_hub_v2.project_level.identity

  identity            = "cmd-probe-template"
  name                = "CMD Probe Template"
  description         = "Command execution probe for custom checks"
  type                = "cmdProbe"
  infrastructure_type = "KubernetesV2"
  tags                = ["cmd", "probe", "custom"]

  # CMD probe configuration
  cmd_probe {
    command = "kubectl get pods -n <+input> | grep Running"
    source  = "inline"
  }

  # Run properties
  run_properties {
    timeout         = "20s"
    interval        = "5s"
    stop_on_failure = false
    verbosity       = "DEBUG"
  }

  # Variables
  variables {
    name        = "check_namespace"
    value       = "<+input>.default('default')"
    type        = "string"
    required    = false
    description = "Namespace to check for running pods"
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
#   - type ("httpProbe", "cmdProbe", "k8sProbe", "promProbe")
#   - infrastructure_type (e.g., "Kubernetes", "KubernetesV2")
#   - One of: http_probe, cmd_probe, k8s_probe, prom_probe
#
# Optional:
#   - description
#   - tags
#   - run_properties (timeout, interval, polling_interval, stop_on_failure, verbosity)
#   - variables (name, value, type, required, description)
#
# Run Properties:
#   - timeout: Maximum time for probe execution
#   - interval: Time between probe executions
#   - polling_interval: Time between status checks
#   - stop_on_failure: Stop experiment if probe fails
#   - verbosity: Log level (INFO, DEBUG, ERROR)
#
# Runtime Inputs:
#   - Use "<+input>" for required runtime input
#   - Use "<+input>.default('value')" for optional with default
#
# Computed (read-only):
#   - probe_id, created_at, updated_at
#
# Import Format: org_id/project_id/hub_identity/identity
# Example: terraform import harness_chaos_probe_template.example \
#          my_org/my_project/my-hub/my-probe
# ----------------------------------------------------------------------------
