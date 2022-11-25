resource "harness_platform_gitops_gnupg" "example" {
  account_id = "account_id"
  agent_id   = "agent_id"

  request {
    upsert = true
    publickey {
      key_data = "-----BEGIN PGP PUBLIC KEY BLOCK-----XXXXXX-----END PGP PUBLIC KEY BLOCK-----"
    }
  }
  lifecycle {
    ignore_changes = [
      request.0.upsert,
    ]
  }
}