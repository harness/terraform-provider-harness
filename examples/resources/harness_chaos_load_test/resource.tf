# K6 toolConfig is flat. Script content is base64.
resource "harness_chaos_load_test" "k6" {
  org_id     = "default"
  project_id = "demo"
  identifier = "checkout_k6"
  name       = "Checkout k6"

  tool_type   = "K6"
  target_type = "kubernetes"

  environment_id    = "prod"
  infrastructure_id = harness_chaos_infrastructure_v2.payments.infra_id

  cleanup_policy = "delete"

  tool_config = jsonencode({
    mode = "script"
    script = {
      content = base64encode("import http from 'k6/http'; export default function () { http.get('https://example.com'); }")
    }
    tunables = {
      targetUsers     = 10
      durationSeconds = 60
    }
  })
}

# Locust toolConfig is nested under locust. It can target a Linux infrastructure.
resource "harness_chaos_load_test" "locust" {
  org_id     = "default"
  project_id = "demo"
  identifier = "browse_locust"
  name       = "Browse locust"

  tool_type   = "Locust"
  target_type = "linux"

  environment_id    = "prod"
  infrastructure_id = "linux-vm-1"
  cleanup_policy    = "retain"

  tool_config = jsonencode({
    locust = {
      mode = "script"
      script = {
        content = base64encode("from locust import HttpUser, task\nclass Browse(HttpUser):\n    @task\n    def home(self):\n        self.client.get('/')\n")
      }
    }
  })
}

# K6 image mode. The container image bundles the script. hostUrl and iterations are tunables.
resource "harness_chaos_load_test" "k6_image" {
  org_id     = "default"
  project_id = "demo"
  identifier = "checkout_k6_image"
  name       = "Checkout k6 image"

  tool_type   = "K6"
  target_type = "kubernetes"

  environment_id    = "prod"
  infrastructure_id = harness_chaos_infrastructure_v2.payments.infra_id

  tool_config = jsonencode({
    mode = "image"
    script = {
      image = "ghcr.io/example/k6:1"
    }
    tunables = {
      hostUrl    = "https://example.com"
      iterations = 100
    }
  })
}

# Locust image mode. envVars and thresholds are optional and live in the same object.
resource "harness_chaos_load_test" "locust_image" {
  org_id     = "default"
  project_id = "demo"
  identifier = "browse_locust_image"
  name       = "Browse locust image"

  tool_type   = "Locust"
  target_type = "linux"

  environment_id    = "prod"
  infrastructure_id = "linux-vm-1"

  tool_config = jsonencode({
    locust = {
      mode = "image"
      script = {
        image = "ghcr.io/example/locust:1"
      }
      tunables = {
        targetUrl = "https://example.com"
      }
    }
  })
}

# Harness-managed Kubernetes infrastructure. Do not set environment or infrastructure ids.
resource "harness_chaos_load_test" "managed" {
  org_id     = "default"
  project_id = "demo"
  identifier = "hosted_k6"
  name       = "Hosted k6"

  tool_type     = "K6"
  target_type   = "kubernetes"
  managed_infra = true

  tool_config = jsonencode({
    mode = "script"
    script = {
      content = base64encode("import http from 'k6/http'; export default function () { http.get('https://example.com'); }")
    }
  })
}
