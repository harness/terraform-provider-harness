# Import an Account level Gitops Repository
terraform import harness_platform_gitops_repository.example <agent_id>/<respository_id>

# Import an Org level Gitops Repository
terraform import harness_platform_gitops_repository.example <organization_id>/<agent_id>/<respository_id>

# Import a Project level Gitops Repository
terraform import harness_platform_gitops_repository.example <organization_id>/<project_id>/<agent_id>/<respository_id>
