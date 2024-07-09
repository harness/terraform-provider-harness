# Import account level jdbc connector 
terraform import harness_platform_connector_jdbc.example <connector_id>

# Import org level jdbc connector 
terraform import harness_platform_connector_jdbc.example <ord_id>/<connector_id>

# Import project level jdbc connector 
terraform import harness_platform_connector_jdbc.example <org_id>/<project_id>/<connector_id>