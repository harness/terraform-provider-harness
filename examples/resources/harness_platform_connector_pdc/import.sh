# Import account level nexus connector 
terraform import harness_platform_connector_nexus.example <connector_id>

# Import org level nexus connector 
terraform import harness_platform_connector_nexus.example <ord_id>/<connector_id>

# Import project level nexus connector 
terraform import harness_platform_connector_nexus.example <org_id>/<project_id>/<connector_id>
