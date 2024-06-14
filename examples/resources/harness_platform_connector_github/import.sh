# Import account level github connector 
terraform import harness_platform_connector_github.example <connector_id>

# Import org level github connector 
terraform import harness_platform_connector_github.example <ord_id>/<connector_id>

# Import project level github connector 
terraform import harness_platform_connector_github.example <org_id>/<project_id>/<connector_id>
