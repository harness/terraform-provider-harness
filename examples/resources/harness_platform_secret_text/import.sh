# Import account level secret text
terraform import harness_platform_secret_text.example <secret_text_id>

# Import org level secret text
terraform import harness_platform_secret_text.example <ord_id>/<secret_text_id>

# Import project level secret text
terraform import harness_platform_secret_text.example <org_id>/<project_id>/<secret_text_id>
