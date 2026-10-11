# Import the mapping of a kind of an account level integration
terraform import harness_platform_idp_integration_kind_mapping.example <integration_id>/<kind>

# Import the mapping of a kind of an org level integration
terraform import harness_platform_idp_integration_kind_mapping.example <org_id>/<integration_id>/<kind>

# Import the mapping of a kind of a project level integration
terraform import harness_platform_idp_integration_kind_mapping.example <org_id>/<project_id>/<integration_id>/<kind>
