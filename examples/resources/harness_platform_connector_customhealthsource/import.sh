# Import account level elasticsearch connector
terraform import harness_platform_connector_customhealthsource.example <connector_id>

# Import org level elasticsearch connector
terraform import harness_platform_connector_customhealthsource.example <ord_id>/<connector_id>

# Import project level elasticsearch connector
terraform import harness_platform_connector_customhealthsource.example <org_id>/<project_id>/<connector_id>
