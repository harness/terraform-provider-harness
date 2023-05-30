# Import account level token
terraform import harness_platform_token <parent_id>/<apikey_id>/<apikey_type>/<token_id>

# Import org level token
terraform import harness_platform_token <org_id>/<parent_id>/<apikey_id>/<apikey_type>/<token_id>

# Import project level token
terraform import harness_platform_token <org_id>/<project_id>/<parent_id>/<apikey_id>/<apikey_type>/<token_id>