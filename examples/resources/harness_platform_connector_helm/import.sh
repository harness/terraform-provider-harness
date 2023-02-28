# Import account level helm connector 
terraform import harness_platform_connector_http_helm.example <connector_id>

# Import org level helm connector 
terraform import harness_platform_connector_http_helm.example <ord_id>/<connector_id>

# Import project level helm connector 
terraform import harness_platform_connector_http_helm.example <org_id>/<project_id>/<connector_id>
