# Import account level datadog connector 
terraform import harness_platform_connector_datadog.example <connector_id>

# Import org level datadog connector 
terraform import harness_platform_connector_datadog.example <ord_id>/<connector_id>

# Import project level datadog connector 
terraform import harness_platform_connector_datadog.example <org_id>/<project_id>/<connector_id>
