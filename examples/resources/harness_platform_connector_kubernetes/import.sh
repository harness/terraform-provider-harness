# Import account level kubernetes connector 
terraform import harness_platform_connector_kubernetes.example <connector_id>

# Import org level kubernetes connector 
terraform import harness_platform_connector_kubernetes.example <ord_id>/<connector_id>

# Import project level kubernetes connector 
terraform import harness_platform_connector_kubernetes.example <org_id>/<project_id>/<connector_id>
