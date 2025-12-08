# Example 1: Cluster Generator
resource "harness_platform_gitops_applicationset" "cluster_generator" {
  org_id     = "default"
  project_id = "projectId"
  agent_id   = "account.agentuseast1"
  upsert     = true
  
  applicationset {
    metadata {
      name      = "cluster-appset"
      namespace = "argocd"
    }
    spec {
      go_template         = true
      go_template_options = ["missingkey=error"]

      generator {
        clusters {
          enabled = true
        }
      }
      
      template {
        metadata {
          name = "{{.name}}-guestbook"
          labels = {
            env = "dev"
          }
        }
        spec {
          project = "default"
          source {
            repo_url        = "https://github.com/argoproj/argocd-example-apps.git"
            path            = "helm-guestbook"
            target_revision = "HEAD"
          }
          destination {
            server    = "{{.url}}"
            namespace = "app-ns-{{.name}}"
          }
        }
      }
    }
  }
}

# Example 2: List Generator
resource "harness_platform_gitops_applicationset" "list_generator" {
  org_id     = "default"
  project_id = "projectId"
  agent_id   = "account.agentuseast1"
  upsert     = true
  
  applicationset {
    metadata {
      name = "list-appset"
    }
    spec {
      go_template         = true
      go_template_options = ["missingkey=error"]
      
      generator {
        list {
          elements = [
            {
              cluster = "engineering-dev"
              url     = "https://kubernetes.default.svc"
            },
            {
              cluster = "engineering-prod"
              url     = "https://kubernetes.prod.svc"
            }
          ]
        }
      }
      
      template {
        metadata {
          name = "{{.cluster}}-guestbook"
        }
        spec {
          project = "default"
          source {
            repo_url        = "https://github.com/argoproj/argocd-example-apps.git"
            path            = "helm-guestbook"
            target_revision = "HEAD"
          }
          destination {
            server    = "{{.url}}"
            namespace = "default"
          }
        }
      }
    }
  }
}

# Example 3: Git Generator with Files
resource "harness_platform_gitops_applicationset" "git_files" {
  org_id     = "default"
  project_id = "projectId"
  agent_id   = "account.agentuseast1"
  upsert     = true
  
  applicationset {
    metadata {
      name = "git-files-appset"
    }
    spec {
      generator {
        git {
          repo_url = "https://github.com/example/config-repo"
          revision = "main"
          
          file {
            path = "apps/*/config.json"
          }
        }
      }
      
      template {
        metadata {
          name = "{{.path.basename}}-app"
        }
        spec {
          project = "default"
          source {
            repo_url        = "https://github.com/example/app-repo"
            path            = "{{.path.path}}"
            target_revision = "main"
          }
          destination {
            server    = "https://kubernetes.default.svc"
            namespace = "{{.path.basename}}"
          }
        }
      }
    }
  }
}

# Example 4: Git Generator with Directories
resource "harness_platform_gitops_applicationset" "git_directories" {
  org_id     = "default"
  project_id = "projectId"
  agent_id   = "account.agentuseast1"
  upsert     = true
  
  applicationset {
    metadata {
      name = "git-directories-appset"
    }
    spec {
      generator {
        git {
          repo_url = "https://github.com/argoproj/argo-cd.git"
          revision = "HEAD"
          
          directory {
            path    = "applicationset/examples/git-generator-directory/cluster-addons/*"
            exclude = false
          }
        }
      }
      
      template {
        metadata {
          name = "{{.path.basename}}-addon"
        }
        spec {
          project = "default"
          source {
            repo_url        = "https://github.com/argoproj/argo-cd.git"
            path            = "{{.path.path}}"
            target_revision = "HEAD"
          }
          destination {
            server    = "https://kubernetes.default.svc"
            namespace = "{{.path.basename}}"
          }
          sync_policy {
            automated {
              prune     = true
              self_heal = true
            }
          }
        }
      }
    }
  }
}
