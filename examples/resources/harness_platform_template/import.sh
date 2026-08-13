# Import account level template (stable version)
terraform import harness_platform_template.example <template_id>

# Import account level template (specific version)
terraform import harness_platform_template.example <template_id>/versions/<version>

# Import org level template (stable version)
terraform import harness_platform_template.example <org_id>/<template_id>

# Import org level template (specific version)
terraform import harness_platform_template.example <org_id>/<template_id>/versions/<version>

# Import project level template (stable version)
terraform import harness_platform_template.example <org_id>/<project_id>/<template_id>

# Import project level template (specific version)
terraform import harness_platform_template.example <org_id>/<project_id>/<template_id>/versions/<version>
