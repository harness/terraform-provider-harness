resource "harness_platform_gitops_repo_cert" "example" {
  account_id = "account_id"
  agent_id   = "agent_id"

  request {
    upsert = true
    certificates {
      metadata {

      }
      items {
        server_name = "serverName"
        cert_type   = "https"
        cert_data   = "yourcertdata"
      }
    }
  }
}
