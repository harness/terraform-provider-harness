# Import account level ccm filter
terraform import harness_platform_ccm_filters.example <filter_id>/<type>

# Import org level ccm filter
terraform import harness_platform_ccm_filters.example <ord_id>/<filter_id>/<type>

# Import project level ccm filter
terraform import harness_platform_ccm_filters.example <org_id>/<project_id>/<filter_id>/<type>
