# Import gitOps applicationset with account level agent, agent id has account prefix #
terraform import harness_platform_gitops_applicationset.example <organization_id>/<project_id>/<agent_id>/<identifier>

# Import gitOps applicationset with org level agent, agent id has org prefix #
terraform import harness_platform_gitops_applicationset.example <organization_id>/<project_id>/<agent_id>/<identifier>

# Import gitOps applicationset with project level agent #
terraform import harness_platform_gitops_applicationset.example <organization_id>/<project_id>/<agent_id>/<identifier>
