# Import account level jenkins connector 
terraform import harness_platform_connector_jenkins.example <connector_id>

# Import org level jenkins connector 
terraform import harness_platform_connector_jenkins.example <ord_id>/<connector_id>

# Import project level jenkins connector 
terraform import harness_platform_connector_jenkins.example <org_id>/<project_id>/<connector_id>
