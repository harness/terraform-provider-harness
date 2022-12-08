# Import account level template filter
terraform import harness_platform_template_filters.example <filter_id>/<type>

# Import org level template filter
terraform import harness_platform_template_filters.example <ord_id>/<filter_id>/<type>

# Import project level template filter
terraform import harness_platform_template_filters.example <org_id>/<project_id>/<filter_id>/<type>
