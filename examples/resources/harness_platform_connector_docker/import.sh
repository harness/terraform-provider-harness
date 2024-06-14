# Import account level docker connector 
terraform import harness_platform_connector_docker.example <connector_id>

# Import org level docker connector 
terraform import harness_platform_connector_docker.example <ord_id>/<connector_id>

# Import project level docker connector 
terraform import harness_platform_connector_docker.example <org_id>/<project_id>/<connector_id>
