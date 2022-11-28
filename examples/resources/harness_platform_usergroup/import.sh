# Import account level user group
terraform import harness_platform_usergroup.example <usergroup_id>

# Import org level user group
terraform import harness_platform_usergroup.example <ord_id>/<usergroup_id>

# Import project level user group
terraform import harness_platform_usergroup.example <org_id>/<project_id>/<usergroup_id>
