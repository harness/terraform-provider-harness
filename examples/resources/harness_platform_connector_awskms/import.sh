# Import account level awskms connector 
terraform import harness_platform_connector_awskms.example <connector_id>

# Import org level awskms connector 
terraform import harness_platform_connector_awskms.example <ord_id>/<connector_id>

# Import project level awskms connector 
terraform import harness_platform_connector_awskms.example <org_id>/<project_id>/<connector_id>
