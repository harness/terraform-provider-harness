# Plugin with multiple env_variables and proxy entries
resource "harness_platform_idp_plugin" "advanced" {
  identifier = "harness-proxy"
  name       = "Configure Backend Proxies"
  enabled    = true
  configs    = file("${path.module}/plugin_configs.yaml")

  env_variables {
    env_name                  = "PRIMARY_KEY"
    type                      = "Secret"
    harness_secret_identifier = "primary_secret"
  }

  env_variables {
    env_name                  = "SECONDARY_KEY"
    type                      = "Secret"
    harness_secret_identifier = "secondary_secret"
  }

  proxy {
    host              = "app.harness.io"
    proxy             = true
    selectors         = ["default"]
    health_check_path = "/health"
  }

  proxy {
    host      = "internal.mycompany.net"
    proxy     = true
    selectors = ["internal-delegate"]
  }
}

# Minimal plugin configuration
resource "harness_platform_idp_plugin" "basic" {
  identifier = "harness-proxy"
  name       = "Configure Backend Proxies"
  configs    = <<-EOT
    proxy:
      endpoints:
        /api:
          target: https://app.harness.io/gateway
    EOT
}

# Plugin with env_variables and proxy configuration
resource "harness_platform_idp_plugin" "full" {
  identifier = "harness-proxy"
  name       = "Configure Backend Proxies"
  enabled    = true
  configs    = <<-EOT
    proxy:
      endpoints:
        /api:
          target: https://app.harness.io/gateway
          pathRewrite:
            api/proxy/?: /
          headers:
            x-api-key: ${API_KEY}
    EOT

  env_variables {
    env_name                  = "API_KEY"
    type                      = "Secret"
    harness_secret_identifier = "my_harness_secret"
  }

  proxy {
    host      = "app.harness.io"
    proxy     = true
    selectors = ["delegate-selector-1"]
  }
}
