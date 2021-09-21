resource "tls_private_key" "harness_deploy_key" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

data "harness_secret_manager" "secret_manager" {
  default = true
}

resource "harness_encrypted_text" "my_secret" {
  name              = "my_secret"
  value             = tls_private_key.harness_deploy_key.private_key_pem
  secret_manager_id = data.harness_secret_manager.secret_manager.id
}

resource "harness_ssh_credential" "ssh_creds" {
  name = "ssh-test"
  
  ssh_authentication {
    port     = 22
    username = "git"
    inline_ssh {
      ssh_key_file_id = harness_encrypted_text.my_secret.id
    }
  }

  // This is a workaround until https://harness.atlassian.net/browse/PL-17967 is resolved
  // The graphql API currently doesn't return all the fields in the query.
  lifecycle {
    ignore_changes = ["ssh_authentication"]
  }
}
