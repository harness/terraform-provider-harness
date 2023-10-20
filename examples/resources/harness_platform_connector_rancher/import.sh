# Import account level rancher connector 
terraform import harness_platform_connector_rancher.example <connector_id>

# Import org level rancher connector 
terraform import harness_platform_connector_rancher.example <ord_id>/<connector_id>

# Import project level rancher connector 
terraform import harness_platform_connector_rancher.example <org_id>/<project_id>/<connector_id>
