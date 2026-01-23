resource "harness_chaos_action_template" "delay_example" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"

  identity    = "delay-action-template"
  name        = "Delay Action Template"
  description = "A simple delay action for chaos experiments"
  type        = "delay"

  infrastructure_type = "Kubernetes"

  delay_action {
    duration = "30s"
  }

  run_properties {
    timeout         = "5m"
    stop_on_failure = true
  }

  tags = ["delay", "chaos", "kubernetes"]
}
