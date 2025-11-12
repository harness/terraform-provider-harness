resource "harness_autostopping_rule_scale_group" "test" {
  name               = "test"
  cloud_connector_id = "test-connector"
  idle_time_mins     = 5
  custom_domains     = ["app.example.com"]
  scale_group {
    id        = "asg-arn"
    name      = "asg-name"
    region    = "us-east-1"
    desired   = 1
    min       = 1
    max       = 2
    on_demand = 1
  }
  http {
    proxy_id = "lb-id"
    routing {
      source_protocol = "http"
      source_port     = 80
      action          = "forward"
      target_protocol = "http"
      target_port     = 80
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