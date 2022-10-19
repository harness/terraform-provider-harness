data "harness_platform_gitops_cluster" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"

  request {
    upsert = false
    cluster {
      server = "server_test"
      name   = "cluster_name"
      config {
        username = "test_username"
        password = "test_password"
        tls_client_config {
          insecure = true
        }
        cluster_connection_type = "USERNAME_PASSWORD"
      }
    }
  }
}
