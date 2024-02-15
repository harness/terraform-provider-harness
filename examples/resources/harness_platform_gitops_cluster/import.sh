# Import an Account level Gitops Cluster
terraform import harness_platform_gitops_cluster.example <agent_id>/<cluster_id>

# Import an Org level Gitops Cluster
terraform import harness_platform_gitops_cluster.example <organization_id>/<agent_id>/<cluster_id>

# Import a Project level Gitops Cluster
terraform import harness_platform_gitops_cluster.example <organization_id>/<project_id>/<agent_id>/<cluster_id>
