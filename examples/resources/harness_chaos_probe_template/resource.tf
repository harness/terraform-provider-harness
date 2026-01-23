# Example 1: HTTP Probe Template
resource "harness_chaos_probe_template" "http_example" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"

  identity    = "http-health-check"
  name        = "HTTP Health Check Probe"
  description = "Checks application health via HTTP endpoint"
  type        = "httpProbe"

  http_probe {
    url = "https://api.example.com/health"
    
    method {
      get {
        criteria      = "=="
        response_code = "200"
      }
    }
  }

  run_properties {
    timeout          = "30s"
    interval         = "5s"
    polling_interval = "2s"
    stop_on_failure  = false
  }

  tags = ["http", "health-check", "api"]
}

# Example 2: Command Probe Template
resource "harness_chaos_probe_template" "cmd_example" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"

  identity    = "disk-usage-check"
  name        = "Disk Usage Check Probe"
  description = "Checks disk usage on target system"
  type        = "cmdProbe"

  cmd_probe {
    command = "df -h | grep '/data' | awk '{print $5}' | sed 's/%//'"
    
    source {
      inline {
        command = "df -h | grep '/data' | awk '{print $5}' | sed 's/%//'"
      }
    }

    comparator {
      type     = "int"
      criteria = "<"
      value    = "80"
    }
  }

  run_properties {
    timeout  = "10s"
    interval = "5s"
  }

  tags = ["command", "disk", "monitoring"]
}

# Example 3: Kubernetes Probe Template
resource "harness_chaos_probe_template" "k8s_example" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"

  identity    = "pod-ready-check"
  name        = "Pod Ready Check Probe"
  description = "Verifies pods are in ready state"
  type        = "k8sProbe"

  k8s_probe {
    version   = "v1"
    resource  = "pods"
    namespace = "production"
    operation = "present"
    
    field_selector = "status.phase=Running"
    label_selector = "app=frontend"
  }

  run_properties {
    timeout  = "30s"
    interval = "10s"
  }

  tags = ["kubernetes", "pod", "health"]
}

# Example 4: APM Probe Template - Prometheus
resource "harness_chaos_probe_template" "prometheus_example" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"

  identity    = "cpu-usage-monitor"
  name        = "CPU Usage Monitor"
  description = "Monitors CPU usage via Prometheus"
  type        = "apmProbe"

  apm_probe {
    apm_type = "Prometheus"

    comparator {
      type     = "float"
      criteria = "<="
      value    = "80.0"
    }

    prometheus_inputs {
      connector_id = "prometheus-connector"
      query        = "avg(rate(container_cpu_usage_seconds_total[5m])) * 100"
    }
  }

  run_properties {
    timeout          = "1m"
    interval         = "15s"
    polling_interval = "5s"
  }

  tags = ["prometheus", "apm", "cpu"]
}

# Example 5: Probe Template with Variables
resource "harness_chaos_probe_template" "with_variables" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"

  identity    = "configurable-http-probe"
  name        = "Configurable HTTP Probe"
  description = "HTTP probe with runtime configurable endpoint"
  type        = "httpProbe"

  variables {
    name        = "TARGET_URL"
    description = "Target URL to probe"
    type        = "string"
    value       = "<+input>"
  }

  variables {
    name        = "EXPECTED_CODE"
    description = "Expected HTTP response code"
    type        = "string"
    value       = "<+input>.default('200')"
  }

  http_probe {
    url = "<+variables.TARGET_URL>"
    
    method {
      get {
        criteria      = "=="
        response_code = "<+variables.EXPECTED_CODE>"
      }
    }
  }

  run_properties {
    timeout  = "30s"
    interval = "5s"
  }

  tags = ["http", "configurable", "runtime-input"]
}
