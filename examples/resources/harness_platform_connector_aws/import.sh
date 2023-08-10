# Import account level aws connector
terraform import harness_platform_connector_aws.example <connector_id>

# Import organization level aws connector
terraform import harness_platform_connector_aws.example <organization_id>/<connector_id>

# Import project level aws connector
terraform import harness_platform_connector_aws.example <organization_id>/<project_id>/<connector_id>
