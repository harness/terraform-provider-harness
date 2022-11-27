# Import account level role assignments
terraform import harness_platform_role_assignments.example <role_assignments_id>

# Import org level role assignments
terraform import harness_platform_role_assignments.example <ord_id>/<role_assignments_id>

# Import project level role assignments
terraform import harness_platform_role_assignments.example <org_id>/<project_id>/<role_assignments_id>
