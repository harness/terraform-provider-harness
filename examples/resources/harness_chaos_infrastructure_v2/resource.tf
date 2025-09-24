// Create a new chaos infrastructure with minimal configuration
resource "harness_chaos_infrastructure_v2" "test" {
  # Required fields
  org_id         = "<org_id>"
  project_id     = "<project_id>"
  environment_id = "<environment_id>"
  infra_id       = "<infra_id>"
  name           = "<name>"
  description    = "<description>"
}

// Create a new chaos infrastructure with full configuration
resource "harness_chaos_infrastructure_v2" "test" {
  # Required fields
  org_id         = "<org_id>"
  project_id     = "<project_id>"
  environment_id = "<environment_id>"
  infra_id       = "<infra_id>"

  # Optional fields
  name = "<name>"
  # Expected to be a numeric group ID
  run_as_group = "<run_as_group>"
  run_as_user  = "<run_as_user>"

  description     = "<description>"
  service_account = "<service_account>"
  namespace       = "<namespace>"
  infra_type      = "<infra_type>"

  ai_enabled           = "<ai_enabled>"
  insecure_skip_verify = "<insecure_skip_verify>"

  # Example of node selector
  node_selector = {
    "kubernetes.io/os"   = "linux"
    "kubernetes.io/arch" = "amd64"
  }

  # Example of labels
  label = {
    "label_key" = "<label_value>"
  }

  # Example of annotations
  annotation = {
    "annotation_key" = "<annotation_value>"
  }

  # Example of volumes
  volumes {
    name       = "test-volume"
    size_limit = "1Gi"
  }

  # Example of volume mounts
  volume_mounts {
    name              = "test-volume"
    mount_path        = "/data"
    read_only         = false
    mount_propagation = "None"
  }

  # Example of environment variables
  env {
    name  = "ENV_VAR"
    value = "test-value"
  }

  # Example of tolerations
  tolerations {
    key                = "key1"
    operator           = "Equal"
    value              = "value1"
    effect             = "NoSchedule"
    toleration_seconds = 3600
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

  image_registry {
    registry_account    = "<registry_account>"
    registry_server     = "<registry_server>"
    secret_name         = "<secret_name>"
    is_override_allowed = true
    is_default          = false
    is_private          = true
    use_custom_images   = true

    custom_images {
      ddcr        = "<ddcr_image>"
      ddcr_fault  = "<ddcr_fault_image>"
      ddcr_lib    = "<ddcr_lib_image>"
      log_watcher = "<log_watcher_image>"
    }
  }
  containers = "<containers>"
}
