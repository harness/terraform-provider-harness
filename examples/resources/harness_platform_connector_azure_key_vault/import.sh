# Import account level azure key vault connector 
terraform import harness_platform_connector_azure_key_vault.example <connector_id>

# Import org level azure key vault connector 
terraform import harness_platform_connector_azure_key_vault.example <ord_id>/<connector_id>

# Import project level azure key vault connector 
terraform import harness_platform_connector_azure_key_vault.example <org_id>/<project_id>/<connector_id>
