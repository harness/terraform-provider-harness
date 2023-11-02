# Import an Account level Gitops Repository Credentials 
terraform import harness_platform_gitops_repo_cred.example <agent_id>/<repocred_id>

# Import an Org level Gitops Repository Credentials 
terraform import harness_platform_gitops_repo_cred.example <organization_id>/<agent_id>/<repocred_id>

# Import a Project level Gitops Repository Credentials 
terraform import harness_platform_gitops_repo_cred.example <organization_id>/<project_id>/<agent_id>/<repocred_id>
