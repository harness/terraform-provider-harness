resource "harness_platform_gitops_repo_cert" "example" {
  account_id = "account_id"
  agent_id   = "agent_id"

  request {
    upsert = true
    certificates {
      metadata {

      }
      items {
        server_name   = "github.com"
        cert_type     = "ssh"
        cert_sub_type = "ecdsa-sha2-nistp256"
        cert_data     = "QUFBQUUyVmpaSE5oTFhOb1lUSXRibWx6ZEhBeU5UWUFBQUFJYm1semRIQXlOVFlBQUFCQkJFbUtTRU5qUUVlek9teGtaTXk3b3BLZ3dGQjlua3Q1WVJyWU1qTnVHNU44N3VSZ2c2Q0xyYm81d0FkVC95NnYwbUtWMFUydzBXWjJZQi8rK1Rwb2NrZz0="
      }
    }
  }
}
