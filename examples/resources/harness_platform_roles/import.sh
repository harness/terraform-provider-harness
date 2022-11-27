# Import account level roles
terraform import harness_platform_roles.example <roles_id>

# Import org level roles
terraform import harness_platform_roles.example <ord_id>/<roles_id>

# Import project level roles
terraform import harness_platform_roles.example <org_id>/<project_id>/<roles_id>
