# Import account level pdc connector 
terraform import harness_platform_connector_pdc.example <connector_id>

# Import org level pdc connector 
terraform import harness_platform_connector_pdc.example <ord_id>/<connector_id>

# Import project level pdc connector 
terraform import harness_platform_connector_pdc.example <org_id>/<project_id>/<connector_id>
