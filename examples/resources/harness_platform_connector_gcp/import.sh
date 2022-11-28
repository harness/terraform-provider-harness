# Import account level gcp connector 
terraform import harness_platform_connector_gcp.example <connector_id>

# Import org level gcp connector 
terraform import harness_platform_connector_gcp.example <ord_id>/<connector_id>

# Import project level gcp connector 
terraform import harness_platform_connector_gcp.example <org_id>/<project_id>/<connector_id>
