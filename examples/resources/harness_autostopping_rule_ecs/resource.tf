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
