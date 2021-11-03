# Import an account level connector
terraform import harness_connector.example <connector_id>

# Import an organization level connector
terraform import harness_connector.example <organization_id>/<connector_id>

# Import an project level connector
terraform import harness_connector.example <organization_id>/<project_id>/<connector_id>
