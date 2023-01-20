# Import an account level freeze
terraform import harness_connector.example <freeze_id>

# Import an organization level freeze
terraform import harness_connector.example <org_id>/<freeze_id>

# Import project level freeze
terraform import harness_platform_manual_freeze.example <org_id>/<project_id>/<freeze_id>
