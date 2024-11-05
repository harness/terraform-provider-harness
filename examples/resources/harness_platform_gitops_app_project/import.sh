# Import gitOps project with account level agent
terraform import harness_platform_gitops_app_project.example <agent_id>/<app_proj_name>

# Import gitOps project with org level agent
terraform import harness_platform_gitops_app_project.example <organization_id>/<agent_id>/<app_proj_name>

# Import gitOps project with project level agent
terraform import harness_platform_gitops_app_project.example <organization_id>/<project_id>/<agent_id>/<app_proj_name>
