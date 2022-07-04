resource "harness_platform_secret_sshkey" "key_tab_file_path" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  kerberos {
    tgt_key_tab_file_path_spec {
      key_path = "key_path"
    }
    principal             = "principal"
    realm                 = "realm"
    tgt_generation_method = "KeyTabFilePath"
  }
}

resource "harness_platform_secret_sshkey" " tgt_password" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  kerberos {
    tgt_password_spec {
      password = "password"
    }
    principal             = "principal"
    realm                 = "realm"
    tgt_generation_method = "Password"
  }
}

resource "harness_platform_secret_sshkey" "sshkey_reference" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  ssh {
    sshkey_reference_credential {
      user_name            = "user_name"
      key                  = "key"
      encrypted_passphrase = "encrypted_passphrase"
    }
    credential_type = "KeyReference"
  }
}

resource "harness_platform_secret_sshkey" " sshkey_path" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  ssh {
    sshkey_path_credential {
      user_name            = "user_name"
      key_path             = "key_path"
      encrypted_passphrase = "encrypted_passphrase"
    }
    credential_type = "KeyPath"
  }
}

resource "harness_platform_secret_sshkey" "ssh_password" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  ssh {
    ssh_password_credential {
      user_name = "user_name"
      password  = "password"
    }
    credential_type = "Password"
  }
}
