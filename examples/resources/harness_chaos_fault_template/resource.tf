# Example 1: Basic Fault Template
resource "harness_chaos_fault_template" "basic_example" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"

  identity    = "pod-delete-fault"
  name        = "Pod Delete Fault"
  description = "Deletes pods to test resilience"
  
  category        = ["Kubernetes"]
  infrastructures = ["KubernetesV2"]
  type            = "Custom"
  tags            = ["kubernetes", "pod", "chaos"]

  spec {
    chaos {
      fault_name = "pod-delete"

      params {
        key   = "TOTAL_CHAOS_DURATION"
        value = "60"
      }

      params {
        key   = "CHAOS_INTERVAL"
        value = "10"
      }
    }
  }
}

# Example 2: Fault Template with Kubernetes Spec
resource "harness_chaos_fault_template" "with_kubernetes_spec" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"

  identity    = "custom-chaos-runner"
  name        = "Custom Chaos Runner"
  description = "Custom fault with specific runner configuration"
  
  category        = ["Kubernetes"]
  infrastructures = ["KubernetesV2"]
  type            = "Custom"

  spec {
    chaos {
      fault_name = "byoc-injector"

      params {
        key   = "TOTAL_CHAOS_DURATION"
        value = "120"
      }

      kubernetes {
        image             = "chaosnative/chaos-go-runner:ci"
        image_pull_policy = "Always"
        command           = ["/bin/sh"]
        args              = ["-c", "echo 'Starting chaos injection'"]

        labels = {
          "app"         = "chaos"
          "environment" = "production"
        }

        annotations = {
          "chaos.io/type" = "custom"
        }

        resources {
          limits = {
            "cpu"    = "500m"
            "memory" = "512Mi"
          }
          requests = {
            "cpu"    = "250m"
            "memory" = "256Mi"
          }
        }

        env {
          name  = "CHAOS_NAMESPACE"
          value = "default"
        }

        env {
          name  = "LOG_LEVEL"
          value = "info"
        }
      }
    }
  }

  tags = ["kubernetes", "custom", "byoc"]
}

# Example 3: Fault Template with Volumes
resource "harness_chaos_fault_template" "with_volumes" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"

  identity    = "fault-with-volumes"
  name        = "Fault with Volumes"
  description = "Fault template with ConfigMap, Secret, and HostPath volumes"
  
  category        = ["Kubernetes"]
  infrastructures = ["KubernetesV2"]
  type            = "Custom"

  spec {
    chaos {
      fault_name = "byoc-injector"

      params {
        key   = "TOTAL_CHAOS_DURATION"
        value = "60"
      }

      kubernetes {
        image             = "chaosnative/chaos-go-runner:ci"
        image_pull_policy = "Always"

        # ConfigMap volume
        config_map_volume {
          name       = "chaos-config"
          mount_path = "/etc/chaos/config"
        }

        # Secret volume
        secret_volume {
          name       = "chaos-secrets"
          mount_path = "/etc/chaos/secrets"
        }

        # HostPath volume
        host_path_volume {
          name       = "host-data"
          mount_path = "/host/data"
          host_path  = "/var/lib/chaos"
          type       = "Directory"
        }
      }
    }
  }

  tags = ["kubernetes", "volumes", "storage"]
}

# Example 4: Fault Template with Targets
resource "harness_chaos_fault_template" "with_targets" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"

  identity    = "targeted-pod-delete"
  name        = "Targeted Pod Delete"
  description = "Pod delete fault with specific target selection"
  
  category        = ["Kubernetes"]
  infrastructures = ["KubernetesV2"]
  type            = "Custom"

  spec {
    chaos {
      fault_name = "pod-delete"

      params {
        key   = "TOTAL_CHAOS_DURATION"
        value = "60"
      }
    }

    target {
      kubernetes {
        kind      = "deployment"
        namespace = "production"
        names     = "frontend-app"
        labels    = "app=frontend,tier=web"
      }
    }
  }

  tags = ["kubernetes", "pod-delete", "targeted"]
}

# Example 5: Fault Template with Variables
resource "harness_chaos_fault_template" "with_variables" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"

  identity    = "configurable-pod-delete"
  name        = "Configurable Pod Delete"
  description = "Pod delete fault with runtime configurable parameters"
  
  category        = ["Kubernetes"]
  infrastructures = ["KubernetesV2"]
  type            = "Custom"

  # Define template variables
  variables {
    name        = "TARGET_NAMESPACE"
    description = "Namespace to target for chaos"
    type        = "string"
    value       = "<+input>"
  }

  variables {
    name        = "CHAOS_DURATION"
    description = "Duration of chaos in seconds"
    type        = "string"
    value       = "<+input>.default('60')"
  }

  variables {
    name        = "POD_AFFECTED_PERCENTAGE"
    description = "Percentage of pods to affect"
    type        = "string"
    value       = "<+input>.default('50')"
  }

  spec {
    chaos {
      fault_name = "pod-delete"

      # Use variables in params
      params {
        key   = "TARGET_NAMESPACE"
        value = "<+variables.TARGET_NAMESPACE>"
      }

      params {
        key   = "TOTAL_CHAOS_DURATION"
        value = "<+variables.CHAOS_DURATION>"
      }

      params {
        key   = "PODS_AFFECTED_PERC"
        value = "<+variables.POD_AFFECTED_PERCENTAGE>"
      }
    }
  }

  tags = ["kubernetes", "configurable", "runtime-input"]
}

# Example 6: Fault Template with Links
resource "harness_chaos_fault_template" "with_links" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"

  identity    = "documented-fault"
  name        = "Well Documented Fault"
  description = "Fault template with documentation links"
  
  category        = ["Kubernetes"]
  infrastructures = ["KubernetesV2"]
  type            = "Custom"

  # Documentation links
  links {
    name = "Official Documentation"
    url  = "https://docs.harness.io/chaos/faults/pod-delete"
  }

  links {
    name = "Troubleshooting Guide"
    url  = "https://docs.harness.io/chaos/troubleshooting"
  }

  links {
    name = "Source Code"
    url  = "https://github.com/harness/chaos-faults/tree/main/pod-delete"
  }

  spec {
    chaos {
      fault_name = "pod-delete"

      params {
        key   = "TOTAL_CHAOS_DURATION"
        value = "60"
      }
    }
  }

  tags = ["kubernetes", "documented", "production-ready"]
}

# Example 7: Comprehensive Fault Template
resource "harness_chaos_fault_template" "comprehensive" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"

  identity    = "comprehensive-chaos-fault"
  name        = "Comprehensive Chaos Fault"
  description = "Full-featured fault template with all options"
  
  category        = ["Kubernetes", "Network"]
  infrastructures = ["KubernetesV2"]
  type            = "Custom"
  
  permissions_required = "Basic"

  # Variables
  variables {
    name        = "TARGET_NAMESPACE"
    description = "Target namespace"
    type        = "string"
    value       = "<+input>"
  }

  # Links
  links {
    name = "Documentation"
    url  = "https://docs.example.com"
  }

  spec {
    chaos {
      fault_name = "byoc-injector"

      params {
        key   = "TOTAL_CHAOS_DURATION"
        value = "120"
      }

      params {
        key   = "TARGET_NAMESPACE"
        value = "<+variables.TARGET_NAMESPACE>"
      }

      kubernetes {
        image             = "chaosnative/chaos-go-runner:ci"
        image_pull_policy = "Always"
        command           = ["/bin/sh"]
        args              = ["-c", "echo 'Chaos injection started'"]

        labels = {
          "app"  = "chaos"
          "type" = "custom"
        }

        annotations = {
          "chaos.io/managed-by" = "terraform"
        }

        resources {
          limits = {
            "cpu"    = "1000m"
            "memory" = "1Gi"
          }
          requests = {
            "cpu"    = "500m"
            "memory" = "512Mi"
          }
        }

        env {
          name  = "CHAOS_NAMESPACE"
          value = "<+variables.TARGET_NAMESPACE>"
        }

        config_map_volume {
          name       = "chaos-config"
          mount_path = "/etc/config"
        }

        secret_volume {
          name       = "chaos-secrets"
          mount_path = "/etc/secrets"
        }

        tolerations {
          key      = "chaos"
          operator = "Equal"
          value    = "true"
          effect   = "NoSchedule"
        }
      }
    }

    target {
      kubernetes {
        kind      = "deployment"
        namespace = "<+variables.TARGET_NAMESPACE>"
        labels    = "app=frontend"
      }
    }
  }

  tags = ["kubernetes", "comprehensive", "production"]
}
