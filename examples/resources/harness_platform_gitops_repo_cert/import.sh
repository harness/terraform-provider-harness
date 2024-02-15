# Import an Account level Gitops Repository Certificate
terraform import harness_platform_gitops_repo_cert.example <repocert_id>

# Import an Org level Gitops Repository Certificate
terraform import harness_platform_gitops_repo_cert.example <organization_id>/<repocert_id>

# Import a Project level Gitops Repository Certificate
terraform import harness_platform_gitops_repo_cert.example <organization_id>/<project_id>/<repocert_id>
