# Import list of account level service overrides using the env id associated with them
terraform import harness_platform_environment_service_overrides.example <env_id>

# Import list of org level service overrides using the env id associated with them
terraform import harness_platform_environment_service_overrides.example <org_id>/<env_id>

# Import list of project level service overrides using the env id associated with them
terraform import harness_platform_environment_service_overrides.example <org_id>/<project_id>/<env_id>
