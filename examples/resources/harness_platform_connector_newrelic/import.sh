# Import account level newrelic connector 
terraform import harness_platform_connector_newrelic.example <connector_id>

# Import org level newrelic connector 
terraform import harness_platform_connector_newrelic.example <ord_id>/<connector_id>

# Import project level newrelic connector 
terraform import harness_platform_connector_newrelic.example <org_id>/<project_id>/<connector_id>
