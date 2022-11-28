# Import account level gcp secret manager connector 
terraform import harness_platform_connector_gcp_secret_manager.example <connector_id>

# Import org level gcp secret manager connector 
terraform import harness_platform_connector_gcp_secret_manager.example <ord_id>/<connector_id>

# Import project level gcp secret manager connector 
terraform import harness_platform_connector_gcp_secret_manager.example <org_id>/<project_id>/<connector_id>
