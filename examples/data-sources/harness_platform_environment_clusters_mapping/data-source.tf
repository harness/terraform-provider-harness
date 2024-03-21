# data source for gitops clusters mapped to a project level env
data "harness_platform_environment_clusters_mapping" "example" {
  identifier = "mycustomidentifier"
  org_id     = "orgIdentifer"
  project_id = "projectIdentifier"
  env_id     = "exampleEnvId"
  clusters {
    identifier       = "incluster"
    name             = "in-cluster"
    agent_identifier = "account.gitopsagentdev"
    scope            = "ACCOUNT"
  }
}


# data source for two gitops clusters mapped to an account level env
data "harness_platform_environment_clusters_mapping" "example2" {
  identifier = "mycustomidentifier"
  env_id     = "env1"
  clusters {
    identifier       = "clusterA"
    name             = "cluster-A"
    agent_identifier = "account.gitopsagentprod"
    scope            = "ACCOUNT"
  }
  clusters {
    identifier       = "clusterB"
    name             = "cluster-B"
    agent_identifier = "account.gitopsagentprod"
    scope            = "ACCOUNT"
  }
}