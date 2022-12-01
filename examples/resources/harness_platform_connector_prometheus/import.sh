# Import account level prometheus connector 
terraform import harness_platform_connector_prometheus.example <connector_id>

# Import org level prometheus connector 
terraform import harness_platform_connector_prometheus.example <ord_id>/<connector_id>

# Import project level prometheus connector 
terraform import harness_platform_connector_prometheus.example <org_id>/<project_id>/<connector_id>
