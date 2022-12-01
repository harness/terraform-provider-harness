# Import account level vault connector 
terraform import harness_platform_connector_vault.example <connector_id>

# Import org level vault connector 
terraform import harness_platform_connector_vault.example <ord_id>/<connector_id>

# Import project level vault connector 
terraform import harness_platform_connector_vault.example <org_id>/<project_id>/<connector_id>
