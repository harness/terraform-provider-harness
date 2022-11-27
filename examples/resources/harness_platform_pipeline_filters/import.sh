# Import account level pipeline filter
terraform import harness_platform_pipeline_filters.example <filter_id>/<type>

# Import org level pipeline filter
terraform import harness_platform_pipeline_filters.example <ord_id>/<filter_id>/<type>

# Import project level pipeline filter
terraform import harness_platform_pipeline_filters.example <org_id>/<project_id>/<filter_id>/<type>
