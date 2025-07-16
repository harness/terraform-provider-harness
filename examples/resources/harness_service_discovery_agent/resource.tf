// Create a new service discovery agent with minimal configuration
resource "harness_service_discovery_agent" "example" {
  name                   = "example-agent"
  org_identifier         = var.org_identifier
  project_identifier     = var.project_identifier
  environment_identifier = var.environment_identifier
  infra_identifier       = "example-infra"

  config {
    kubernetes {
      namespace = "harness-sd"
    }
  }
}

// Create a new service discovery agent with node agent enabled
resource "harness_service_discovery_agent" "node_agent" {
  name                   = "node-agent-example"
  org_identifier         = var.org_identifier
  project_identifier     = var.project_identifier
  environment_identifier = var.environment_identifier
  infra_identifier       = "node-agent-example"

  config {
    kubernetes {
      namespace = "harness-sd"
    }

    data {
      enable_node_agent = true
    }
  }
}


// Create a new service discovery agent with full configuration
resource "harness_service_discovery_agent" "full_config" {
  name                   = "full-config-example"
  description            = "Example service discovery agent with full configuration"
  org_identifier         = var.org_identifier
  project_identifier     = var.project_identifier
  environment_identifier = var.environment_identifier
  infra_identifier       = "full-config-example"
  permanent_installation = false
  correlation_id         = "full-config-correlation-123"

  tags = {
    "managed-by"  = "terraform"
    "environment" = "dev"
  }

  config {
    collector_image    = "harness/service-discovery-collector:main-latest"
    log_watcher_image  = "harness/chaos-log-watcher:main-latest"
    skip_secure_verify = false

    kubernetes {
      namespace         = "harness-sd"
      service_account   = "harness-sd-sa"
      image_pull_policy = "IfNotPresent"
      run_as_user       = 2000
      run_as_group      = 2000

      labels = {
        "app" = "service-discovery"
        "env" = "dev"
      }

      annotations = {
        "example.com/annotation" = "value"
      }

      node_selector = {
        "kubernetes.io/os" = "linux"
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

      tolerations {
        key      = "key1"
        operator = "Equal"
        value    = "value1"
        effect   = "NoSchedule"
      }
    }

    data {
      enable_node_agent        = true
      node_agent_selector      = "node-role.kubernetes.io/worker="
      enable_batch_resources   = true
      enable_orphaned_pod      = true
      namespace_selector       = "environment=dev"
      collection_window_in_min = 15

      blacklisted_namespaces = ["kube-system", "kube-public"]
      observed_namespaces    = ["default", "harness"]

      cron {
        expression = "0/10 * * * *"
      }
    }

    mtls {
      cert_path   = "/etc/certs/tls.crt"
      key_path    = "/etc/certs/tls.key"
      secret_name = "mtls-secret"
      url         = "https://mtls.example.com:8443"
    }

    proxy {
      http_proxy  = "http://proxy.example.com:8080"
      https_proxy = "https://proxy.example.com:8080"
      no_proxy    = "localhost,127.0.0.1,.svc,.cluster.local"
      url         = "https://proxy.example.com"
    }
  }
}
