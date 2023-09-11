# Import account level environment group.
terraform import harness_platform_environment_group.example <environment_group_id>

# Import org level environment group.
terraform import harness_platform_environment_group.example <org_id>/<environment_group_id>

# Import project level environment group.
terraform import harness_platform_environment_group.example <org_id>/<project_id>/<environment_group_id>
