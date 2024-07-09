// Create Account level gitOps project
resource "harness_platform_gitops_project" "test" {
  account_id = "accountIdentifier"
  agent_id   = "agentIdentifier"
  upsert     = true
  project {
    metadata {
      generation = "generation"
      name       = "name"
      namespace  = "namespace"
    }
    spec {
      cluster_resource_whitelist {
        group = "*"
        kind  = "*"
      }
      destinations {
        namespace = "*"
        server    = "*"
      }
      source_repos = ["*"]
    }
  }
}

// Create Org level gitOps project
resource "harness_platform_gitops_project" "test" {
  account_id = "accountIdentifier"
  org_id     = "orgIdentifier"
  agent_id   = "agentIdentifier"
  upsert     = true
  project {
    metadata {
      generation = "generation"
      name       = "name"
      namespace  = "namespace"
    }
    spec {
      cluster_resource_whitelist {
        group = "*"
        kind  = "*"
      }
      destinations {
        namespace = "*"
        server    = "*"
      }
      source_repos = ["*"]
    }
  }
}

// Create Project level gitOps project
resource "harness_platform_gitops_project" "test" {
  account_id = "accountIdentifier"
  org_id     = "orgIdentifier"
  agent_id   = "agentIdentifier"
  project_id = "projectIdentifier"
  upsert     = true
  project {
    metadata {
      generation = "generation"
      name       = "name"
      namespace  = "namespace"
    }
    spec {
      cluster_resource_whitelist {
        group = "*"
        kind  = "*"
      }
      destinations {
        namespace = "*"
        server    = "*"
      }
      source_repos = ["*"]
    }
  }
}                        