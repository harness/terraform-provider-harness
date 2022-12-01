# Import account level gitlab connector 
terraform import harness_platform_connector_gitlab.example <connector_id>

# Import org level gitlab connector 
terraform import harness_platform_connector_gitlab.example <ord_id>/<connector_id>

# Import project level gitlab connector 
terraform import harness_platform_connector_gitlab.example <org_id>/<project_id>/<connector_id>
