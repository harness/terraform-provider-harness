resource "harness_autostopping_rule_ecs" "test" {
  name               = "name"
  cloud_connector_id = "cloud_connector_id"
  idle_time_mins     = 10
  dry_run            = true
  container {
    cluster    = "cluster"
    service    = "service"
    region     = "us-east-1"
    task_count = 1
  }
  depends {
    rule_id      = 24576
    delay_in_sec = 5
  }
}

# Cross-account AutoStopping rule: the ECS service is in a different cloud account
# than the proxy (access point). Use proxy_cloud_connector_id to specify the
# cloud connector that owns the proxy.
resource "harness_autostopping_rule_ecs" "cross_account" {
  name               = "cross-account-ecs-rule"
  cloud_connector_id = "target_account_connector_id"
  idle_time_mins     = 10
  dry_run            = true
  container {
    cluster    = "cluster"
    service    = "service"
    region     = "us-east-1"
    task_count = 1
  }
  http {
    proxy_id                 = "proxy_id"
    proxy_cloud_connector_id = "proxy_account_connector_id"
    routing {
      source_protocol = "http"
      target_protocol = "http"
      source_port     = 80
      target_port     = 80
      action          = "forward"
    }
    health {
      protocol         = "http"
      port             = 80
      path             = "/"
      timeout          = 30
      status_code_from = 200
      status_code_to   = 299
    }
  }
}
