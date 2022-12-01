# Import account level dynatrace connector 
terraform import harness_platform_connector_dynatrace.example <connector_id>

# Import org level dynatrace connector 
terraform import harness_platform_connector_dynatrace.example <ord_id>/<connector_id>

# Import project level dynatrace connector 
terraform import harness_platform_connector_dynatrace.example <org_id>/<project_id>/<connector_id>
