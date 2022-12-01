# Import account level aws secret manager connector 
terraform import harness_platform_connector_aws_secret_manager.example <connector_id>

# Import org level aws secret manager connector 
terraform import harness_platform_connector_aws_secret_manager.example <ord_id>/<connector_id>

# Import project level aws secret manager connector 
terraform import harness_platform_connector_aws_secret_manager.example <org_id>/<project_id>/<connector_id>
