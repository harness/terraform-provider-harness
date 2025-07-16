# Import account level Service Discovery Agent
terraform import harness_service_discovery_agent.example <environment_identifier>/<infra_identifier>

# Import organization level Service Discovery Agent
terraform import harness_service_discovery_agent.example <org_identifier>/<environment_identifier>/<infra_identifier>

# Import project level Service Discovery Agent
terraform import harness_service_discovery_agent.example <org_identifier>/<project_identifier>/<environment_identifier>/<infra_identifier>
