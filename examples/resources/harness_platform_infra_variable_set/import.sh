# Import account level variable set
terraform import harness_platform_infra_variable_set.example <variable_set_id>

# Import org level variable set
terraform import harness_platform_infra_variable_set.example <ord_id>/<variable_set_id>

# Import project level variable set
terraform import harness_platform_secret_text.example <org_id>/<project_id>/<variable_set_id>

