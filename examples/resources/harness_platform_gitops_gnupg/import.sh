# Import a Account level Gitops Cluster
terraform import harness_platform_gitops_gnupg.example <agent_id>/<key_id>

# Import a Project level Gitops Cluster
terraform import harness_platform_gitops_gnupg.example <organization_id>/<project_id>/<agent_id>/<key_id>