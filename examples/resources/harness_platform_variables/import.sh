# Import account level variables
terraform import harness_platform_variables.example <variable_id>

# Import org level variables
terraform import harness_platform_variables.example <ord_id>/<variable_id>

# Import project level variables
terraform import harness_platform_variables.example <org_id>/<project_id>/<variable_id>
