data "harness_chaos_action_template" "by_identity" {
  org_id       = "my_org"
  project_id   = "my_project"
  hub_identity = "my-chaos-hub"
  identity     = "delay-action-template"
}

output "template_id" {
  value = data.harness_chaos_action_template.by_identity.id
}
