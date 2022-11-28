# Import account level secret file
terraform import harness_platform_secret_file.example <secret_file_id>

# Import org level secret file
terraform import harness_platform_secret_file.example <ord_id>/<secret_file_id>

# Import project level secret file
terraform import harness_platform_secret_file.example <org_id>/<project_id>/<secret_file_id>
