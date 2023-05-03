# Import account level tas connector
terraform import harness_platform_connector_tas.example <connector_id>

# Import organization level tas connector
terraform import harness_platform_connector_tas.example <organization_id>/<connector_id>

# Import project level tas connector
terraform import harness_platform_connector_tas.example <organization_id>/<project_id>/<connector_id>
