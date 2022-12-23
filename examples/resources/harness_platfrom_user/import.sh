# Import account level
terraform import harness_platform_user.john_doe <email_id>

# Import org level 
terraform import harness_platform_user.john_doe <ord_id>/<email_id>

# Import project level
terraform import harness_platform_user.john_doe <org_id>/<project_id>/<email_id>
