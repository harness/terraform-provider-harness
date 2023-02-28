# Import account level infrastructure
terraform import harness_platform_infrastructure.example <env_id>/<infrastructure_id>

# Import org level infrastructure
terraform import harness_platform_infrastructure.example <org_id>/<env_id>/<infrastructure_id>

# Import project level infrastructure
terraform import harness_platform_infrastructure.example <org_id>/<project_id>/<env_id>/<infrastructure_id>
