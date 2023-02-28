# Import account level
terraform import harness_platform_user.john_doe <email_id>

# Import org level 
terraform import harness_platform_user.john_doe <email_id>/<org_id>/

# Import project level
terraform import harness_platform_user.john_doe <email_id>/<org_id>/<project_id>
