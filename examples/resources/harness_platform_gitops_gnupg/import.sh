# Import an Account level Gitops GnuPG Key
terraform import harness_platform_gitops_gnupg.example <agent_id>/<key_id>

# Import an Org level Gitops GnuPG Key
terraform import harness_platform_gitops_gnupg.example <organization_id>/<agent_id>/<key_id>

# Import a Project level Gitops GnuPG Key
terraform import harness_platform_gitops_gnupg.example <organization_id>/<project_id>/<agent_id>/<key_id>
