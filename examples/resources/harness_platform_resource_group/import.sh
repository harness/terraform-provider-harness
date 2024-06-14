# Import account level resource group
terraform import harness_platform_resource_group.example <resource_group_id>

# Import org level resource group
terraform import harness_platform_resource_group.example <ord_id>/<resource_group_id>

# Import project level resource group
terraform import harness_platform_resource_group.example <org_id>/<project_id>/<resource_group_id>
