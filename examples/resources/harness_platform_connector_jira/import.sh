# Import account level jira connector 
terraform import harness_platform_connector_jira.example <connector_id>

# Import org level jira connector 
terraform import harness_platform_connector_jira.example <ord_id>/<connector_id>

# Import project level jira connector 
terraform import harness_platform_connector_jira.example <org_id>/<project_id>/<connector_id>
