# Import account level artifactory connector 
terraform import harness_platform_connector_artifactory.example <connector_id>

# Import org level artifactory connector 
terraform import harness_platform_connector_artifactory.example <ord_id>/<connector_id>

# Import project level artifactory connector 
terraform import harness_platform_connector_artifactory.example <org_id>/<project_id>/<connector_id>
