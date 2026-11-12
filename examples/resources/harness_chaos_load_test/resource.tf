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

# REFERENCE import. The template stays the source of the script, variables, and tool config.
# Later template edits apply to this load test. Do not set tool_config, variable, or max_duration_sec.
# revision is a real revision id. latest is rejected on create.
# tool_type must match the template. managed_infra is rejected: environment_id and infrastructure_id are required.
resource "harness_chaos_load_test" "reference" {
  org_id     = "default"
  project_id = "demo"
  identifier = "checkout_from_template"
  name       = "Checkout from template"

  tool_type   = "K6"
  target_type = "kubernetes"
  import_type = "REFERENCE"

  environment_id    = "prod"
  infrastructure_id = harness_chaos_infrastructure_v2.payments.infra_id

  template_reference {
    hub_identity = "payments-hub"
    identity     = "checkout_template"
    revision     = "1"
  }
}

# LOCAL import. This is a copy. Later template edits do not change this load test.
# tool_config and variable replace the copy. Leave them unset to keep the template's values.
resource "harness_chaos_load_test" "copy" {
  org_id     = "default"
  project_id = "demo"
  identifier = "checkout_copy"
  name       = "Checkout copy"

  tool_type   = "K6"
  target_type = "kubernetes"
  import_type = "LOCAL"

  environment_id    = "prod"
  infrastructure_id = harness_chaos_infrastructure_v2.payments.infra_id

  template_reference {
    hub_identity = "payments-hub"
    identity     = "checkout_template"
    revision     = "1"
  }

  tool_config = jsonencode({
    mode = "script"
    script = {
      content = base64encode("import http from 'k6/http'; export default function () { http.get('https://example.com'); }")
    }
  })
}

# Account hub: leave organization_id and project_id empty.
# Organization hub: set organization_id and leave project_id empty.
resource "harness_chaos_load_test" "account_hub" {
  org_id     = "default"
  project_id = "demo"
  identifier = "checkout_account_template"
  name       = "Checkout from the account hub"

  tool_type   = "Locust"
  import_type = "REFERENCE"

  environment_id    = "prod"
  infrastructure_id = "linux-vm-1"

  template_reference {
    hub_identity = "shared-hub"
    identity     = "browse_template"
    revision     = "1"
  }
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
