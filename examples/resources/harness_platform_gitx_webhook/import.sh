# Import an Account level Gitops Repository
terraform import harness_platform_gitx_webhook.example <webhook_identifier>

# Import an Org level Gitops Repository
terraform import harness_platform_gitx_webhook.example <webhook_identifier>/<org_id>/

# Import a Project level Gitops Repository
terraform import harness_platform_gitx_webhook.example <webhook_identifier>/<org_id>/<project_id>
