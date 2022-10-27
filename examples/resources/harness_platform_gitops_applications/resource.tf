resource "harness_platform_gitops_applications" "example" {
  application {
    metadata {
      annotations = {}
      labels = {
        "harness.io/serviceRef" = "service_id"
        "harness.io/envRef" = "env_id"
      }
      name = "appname123"
    }
    spec {
      sync_policy {
        sync_options = [
          "PrunePropagationPolicy=undefined",
          "CreateNamespace=false",
          "Validate=false",
          "skipSchemaValidations=false",
          "autoCreateNamespace=false",
          "pruneLast=false",
          "applyOutofSyncOnly=false",
          "Replace=false",
          "retry=false"
        ]
      }
      source {
        target_revision = "master"
        repo_url = "https://github.com/willycoll/argocd-example-apps.git"
        path = "helm-guestbook"

      }
      destination {
        namespace = "namespace-123"
        server = "https://1.3.4.5"
      }
    }
  }
  project_id = "project_id"
  org_id = "org_id"
  account_id = "account_id"
  identifier = "identifier"
  cluster_id = "cluster_id"
  repo_id = "repo_id"
  agent_id = "agent_id"
}