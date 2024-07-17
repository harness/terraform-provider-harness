resource "harness_platform_secret_wirm" "key_tab_file_path" {
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

resource "harness_platform_secret_winrm" "tgt_password" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  useSSL = true
  skipCert = true
  useNoProfile = true

  kerberos {
    tgt_password_spec {
      password = "account.${secret.id}"
    }
    principal             = "principal"
    realm                 = "realm"
    tgt_generation_method = "Password"
  }
  parameters = ["key:value", "key2:value2"]
}

resource "harness_platform_secret_wirm" "wirm_path" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  ssh {
    wirm_path_credential {
      user_name            = "user_name"
      key_path             = "key_path"
      encrypted_passphrase = "encrypted_passphrase"
    }
    credential_type = "KeyPath"
  }
}


resource "harness_platform_secret_wirm" "ssh_password" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  ssh {
    ssh_password_credential {
      user_name = "user_name"
      password  = "account.${secret.id}"
    }
    credential_type = "Password"
  }
}

resource "harness_platform_secret_wirm" "ntlm" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 5985
  useSSL = true
  skipCert = true
  useNoProfile = true
  ntlm {
    domain         = "DOMAIN"
    username       = "USERNAME"
    password       = "account.${secret.id}"
  }
  parameters = ["key:value", "key2:value2"]
}
