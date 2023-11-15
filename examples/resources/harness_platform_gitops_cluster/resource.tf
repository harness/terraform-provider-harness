# Clusters without Optional tags
resource "harness_platform_gitops_cluster" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"

  request {
    upsert = false
    cluster {
      server = "https://kubernetes.default.svc"
      name   = "name"
      config {
        tls_client_config {
          insecure = true
        }
        cluster_connection_type = "IN_CLUSTER"
      }

    }
  }
  lifecycle {
    ignore_changes = [
      request.0.upsert, request.0.cluster.0.config.0.bearer_token,
    ]
  }
}

# Clusters with Optional tags
resource "harness_platform_gitops_cluster" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"

  request {
    upsert = false
    tags = [
      "foo:bar",
    ]
    cluster {
      server = "https://kubernetes.default.svc"
      name   = "name"
      config {
        tls_client_config {
          insecure = true
        }
        cluster_connection_type = "IN_CLUSTER"
      }

    }
  }
  lifecycle {
    ignore_changes = [
      request.0.upsert, request.0.cluster.0.config.0.bearer_token,
    ]
  }
}


# Clusters with self signed certificate
resource "harness_platform_gitops_cluster" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"

  request {
    upsert = false
    cluster {
      server = "https://1.2.3.4"
      name   = "name"
      config {
        tls_client_config {
          bearer_token = "ey......X"
          ca_data      = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tClhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWApYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFgKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQ=="
        }
        cluster_connection_type = "SERVICE_ACCOUNT"
      }

    }
  }
  lifecycle {
    ignore_changes = [
      request.0.upsert, request.0.cluster.0.config.0.bearer_token,
    ]
  }
}