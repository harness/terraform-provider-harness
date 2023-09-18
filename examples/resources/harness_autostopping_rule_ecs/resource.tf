resource "harness_autostopping_rule_ecs" "test" {
  name               = "name"
  cloud_connector_id = "cloud_connector_id"
  idle_time_mins     = 10
  container {
    cluster    = "cluster"
    service    = "service"
    region     = "us-east-1"
    task_count = 1
  }
  tcp {
    proxy_id = "proxy_id"
    forward_rule {
      port = 2233
    }
  }
  depends {
    rule_id      = 24576
    delay_in_sec = 5
  }
}
