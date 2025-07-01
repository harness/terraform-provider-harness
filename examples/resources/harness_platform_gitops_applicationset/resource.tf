resource "harness_platform_gitops_applicationset" "test_fixed" {
  #     THIS APPSET IS A VERY BASIC APPSET TO GENERATE ONE APPLICATION PER CLUSTER THAT'S ATTACHED TO THIS AGENT
  org_id     = "default"
  project_id = "projectId"
  agent_id   = "account.agentuseast1"
  upsert     = true
  applicationset {
    metadata {
      name      = "tf-appset"
      namespace = "argocd"
    }
    spec {
      go_template = true

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
            "harness.io/serviceRef" = "svc1"
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
