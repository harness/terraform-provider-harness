// Create Account level gitOps project
resource "harness_platform_gitops_project" "test" {
  account_id = "account_id"
  agent_id   = "agent_id"
  upsert     = true
  project {
    metadata {
      name          = "name"
      namespace     = "rollouts"
      finalizers    = ["resources-finalizer.argocd.argoproj.io"]
      generate_name = "generate_name"
      labels = {
        v1 = "k1"
      }
      annotations = {
        v1 = "k1"
      }
    }
    spec {
      cluster_resource_whitelist {
        group = "*"
        kind  = "Namespace"
      }
      destinations {
        namespace = "guestbook"
        server    = "https://kubernetes.default.svc"
        name      = "in-cluster"
      }
      roles {
        name        = "read-only"
        description = "Read-only privileges to my-project"
        policies    = ["p, proj:%[3]s:read-only, applications, get, %[3]s/*, allow"]
        groups      = ["my-oidc-group"]
      }
      roles {
        name        = "ci-role"
        description = "Sync privileges for guestbook-dev"
        policies    = ["p, proj:%[3]s:ci-role, applications, sync, %[3]s/guestbook-dev, allow"]
        jwt_tokens {
          iat = "1535390316"
        }
      }
      sync_windows {
        kind         = "allow"
        schedule     = "10 1 * * *"
        duration     = "1h"
        applications = ["*-prod"]
        manual_sync  = "true"
      }
      sync_windows {
        kind       = "deny"
        schedule   = "0 22 * * *"
        duration   = "1h"
        namespaces = ["default"]
      }
      sync_windows {
        kind     = "allow"
        schedule = "0 23 * * *"
        duration = "1h"
        clusters = ["in-cluster", "cluster1"]
      }
      namespace_resource_blacklist {
        group = "group"
        kind  = "ResourceQuota"
      }
      namespace_resource_blacklist {
        group = "group2"
        kind  = "LimitRange"
      }
      namespace_resource_blacklist {
        group = "group3"
        kind  = "NetworkPolicy"
      }
      namespace_resource_whitelist {
        group = "apps"
        kind  = "Deployment"
      }
      namespace_resource_whitelist {
        group = "apps"
        kind  = "StatefulSet"
      }
      orphaned_resources {
        warn = "false"
      }
      source_repos = ["*"]
    }
  }
}

// Create Project level gitOps project
resource "harness_platform_gitops_project" "test" {
  account_id = "account_id"
  agent_id   = "agent_id"
  upsert     = true
  project_id = "project_id"
  org_id     = "org_id"
  project {
    metadata {
      name          = "name"
      namespace     = "rollouts"
      finalizers    = ["resources-finalizer.argocd.argoproj.io"]
      generate_name = "generate_name"
      labels = {
        v1 = "k1"
      }
      annotations = {
        v1 = "k1"
      }
    }
    spec {
      cluster_resource_whitelist {
        group = "*"
        kind  = "Namespace"
      }
      destinations {
        namespace = "guestbook"
        server    = "https://kubernetes.default.svc"
        name      = "in-cluster"
      }
      roles {
        name        = "read-only"
        description = "Read-only privileges to my-project"
        policies    = ["p, proj:%[3]s:read-only, applications, get, %[3]s/*, allow"]
        groups      = ["my-oidc-group"]
      }
      roles {
        name        = "ci-role"
        description = "Sync privileges for guestbook-dev"
        policies    = ["p, proj:%[3]s:ci-role, applications, sync, %[3]s/guestbook-dev, allow"]
        jwt_tokens {
          iat = "1535390316"
        }
      }
      sync_windows {
        kind         = "allow"
        schedule     = "10 1 * * *"
        duration     = "1h"
        applications = ["*-prod"]
        manual_sync  = "true"
      }
      sync_windows {
        kind       = "deny"
        schedule   = "0 22 * * *"
        duration   = "1h"
        namespaces = ["default"]
      }
      sync_windows {
        kind     = "allow"
        schedule = "0 23 * * *"
        duration = "1h"
        clusters = ["in-cluster", "cluster1"]
      }
      namespace_resource_blacklist {
        group = "group"
        kind  = "ResourceQuota"
      }
      namespace_resource_blacklist {
        group = "group2"
        kind  = "LimitRange"
      }
      namespace_resource_blacklist {
        group = "group3"
        kind  = "NetworkPolicy"
      }
      namespace_resource_whitelist {
        group = "apps"
        kind  = "Deployment"
      }
      namespace_resource_whitelist {
        group = "apps"
        kind  = "StatefulSet"
      }
      orphaned_resources {
        warn = "false"
      }
      source_repos = ["*"]
    }
  }
}