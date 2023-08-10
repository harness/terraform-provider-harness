# Import account level elasticsearch connector
terraform import harness_platform_connector_elasticsearch.example <connector_id>

# Import org level elasticsearch connector
terraform import harness_platform_connector_elasticsearch.example <ord_id>/<connector_id>

# Import project level elasticsearch connector
terraform import harness_platform_connector_elasticsearch.example <org_id>/<project_id>/<connector_id>
