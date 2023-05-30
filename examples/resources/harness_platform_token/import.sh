# Import account level token
terraform import harness_platform_token <apikey_id>/<token_id>

# Import org level token
terraform import harness_platform_token <org_id>/<apikey_id>/<token_id>

# Import project level token
terraform import harness_platform_token <org_id>/<project_id>/<apikey_id>/<token_id>