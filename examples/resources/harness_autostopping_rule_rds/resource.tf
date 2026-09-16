resource "harness_autostopping_rule_rds" "test" {
  name               = "name"
  cloud_connector_id = "cloud_connector_id"
  idle_time_mins     = 10
  dry_run            = true
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

# Cross-account AutoStopping rule: the RDS instance is in a different cloud account
# than the proxy (access point). Use proxy_cloud_connector_id to specify the
# cloud connector that owns the proxy.
resource "harness_autostopping_rule_rds" "cross_account" {
  name               = "cross-account-rds-rule"
  cloud_connector_id = "target_account_connector_id"
  idle_time_mins     = 10
  dry_run            = true
  database {
    id     = "database_id"
    region = "us-east-1"
  }
  tcp {
    proxy_id                 = "proxy_id"
    proxy_cloud_connector_id = "proxy_account_connector_id"
    forward_rule {
      port = 2233
    }
  }
}
