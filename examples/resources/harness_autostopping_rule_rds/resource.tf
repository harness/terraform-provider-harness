resource "harness_autostopping_rule_rds" "test" {
  name               = "name"
  cloud_connector_id = "cloud_connector_id"
  idle_time_mins     = 10
  database {
    id     = "database_id"
    region = "region"
  }
  tcp {
    proxy_id = "proxy_id"
    forward_rule {
      port = 2233
    }
  }
}
