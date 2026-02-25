resource "harness_autostopping_rule_k8s" "sujeesh-alert" {
  name               = "demo-rule"
  cloud_connector_id = "gcp_qa"
  k8s_connector_id = "app_cluster"
  k8s_namespace = "default"
  idle_time_mins     = 10
  dry_run            = false
  rule_yaml = file("${path.module}/rule.yaml")
}