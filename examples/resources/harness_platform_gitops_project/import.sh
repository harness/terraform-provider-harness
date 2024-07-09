# Import an Account level Gitops Repository
terraform import harness_platform_gitops_project.example <agent_id>/<query_name>

# Import an Org level Gitops Repository
terraform import harness_platform_gitops_repository.example <organization_id>/<agent_id>/<query_name

# Import a Project level Gitops Repository
terraform import harness_platform_gitops_repository.example <organization_id>/<project_id>/<agent_id>/<query_name>
