# Import account level git connector 
terraform import harness_platform_connector_git.example <connector_id>

# Import org level git connector 
terraform import harness_platform_connector_git.example <ord_id>/<connector_id>

# Import project level git connector 
terraform import harness_platform_connector_git.example <org_id>/<project_id>/<connector_id>
