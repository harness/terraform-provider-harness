# Import account level filter
terraform import harness_platform_filters.example <filter_id>/<type>

# Import org level filter
terraform import harness_platform_filters.example <ord_id>/<filter_id>/<type>

# Import project level filter
terraform import harness_platform_filters.example <org_id>/<project_id>/<filter_id>/<type>
