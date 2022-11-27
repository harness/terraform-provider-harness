# Import account level bitbucket connector 
terraform import harness_platform_connector_bitbucket.example <connector_id>

# Import org level bitbucket connector 
terraform import harness_platform_connector_bitbucket.example <ord_id>/<connector_id>

# Import project level bitbucket connector 
terraform import harness_platform_connector_bitbucket.example <org_id>/<project_id>/<connector_id>
