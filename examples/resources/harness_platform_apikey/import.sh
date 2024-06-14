# Import account level apikey
terraform import harness_platform_apikey <parent_id>/<apikey_id>/<apikey_type>

# Import org level apikey
terraform import harness_platform_apikey <org_id>/<parent_id>/<apikey_id>/<apikey_type>

# Import project level apikey
terraform import harness_platform_apikey <org_id>/<project_id>/<parent_id>/<apikey_id>/<apikey_type>
