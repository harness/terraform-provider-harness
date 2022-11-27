# Import account level pagerduty connector 
terraform import harness_platform_connector_pagerduty.example <connector_id>

# Import org level pagerduty connector 
terraform import harness_platform_connector_pagerduty.example <ord_id>/<connector_id>

# Import project level pagerduty connector 
terraform import harness_platform_connector_pagerduty.example <org_id>/<project_id>/<connector_id>
