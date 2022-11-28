# Import account level sumologic connector 
terraform import harness_platform_connector_sumologic.example <connector_id>

# Import org level sumologic connector 
terraform import harness_platform_connector_sumologic.example <ord_id>/<connector_id>

# Import project level sumologic connector 
terraform import harness_platform_connector_sumologic.example <org_id>/<project_id>/<connector_id>
