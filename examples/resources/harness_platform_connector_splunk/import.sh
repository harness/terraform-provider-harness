# Import account level splunk connector 
terraform import harness_platform_connector_splunk.example <connector_id>

# Import org level splunk connector 
terraform import harness_platform_connector_splunk.example <ord_id>/<connector_id>

# Import project level splunk connector 
terraform import harness_platform_connector_splunk.example <org_id>/<project_id>/<connector_id>
