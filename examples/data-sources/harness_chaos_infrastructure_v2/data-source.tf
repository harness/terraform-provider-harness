// Fetch an existing chaos infrastructure V2 by its identifiers
data "harness_chaos_infrastructure_v2" "example" {
  org_id         = "<org_id>"
  project_id     = "<project_id>"
  environment_id = "<environment_id>"
  infra_id       = "<infra_id>"
}

// The data source exposes the pod resource requirements, autopilot mode, and
// the associated discovery agent, alongside the rest of the infrastructure
// attributes.
output "chaos_infra_resources" {
  value = data.harness_chaos_infrastructure_v2.example.resources
}

output "chaos_infra_autopilot_enabled" {
  value = data.harness_chaos_infrastructure_v2.example.autopilot_enabled
}

output "chaos_infra_discovery_agent_id" {
  value = data.harness_chaos_infrastructure_v2.example.discovery_agent_id
}
