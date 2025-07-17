// Data source to fetch a specific agent by name
data "harness_service_discovery_agent" "by_name" {
  name                   = "example-agent"
  org_identifier         = var.org_identifier
  project_identifier     = var.project_identifier
  environment_identifier = var.environment_identifier
}

output "agent_details_by_name" {
  value = data.harness_service_discovery_agent.by_name
}

// Data source to fetch a specific agent by identity
data "harness_service_discovery_agent" "by_identity" {
  identity               = "example-infra"
  org_identifier         = var.org_identifier
  project_identifier     = var.project_identifier
  environment_identifier = var.environment_identifier
}

output "agent_details_by_identity" {
  value = data.harness_service_discovery_agent.by_identity
}
