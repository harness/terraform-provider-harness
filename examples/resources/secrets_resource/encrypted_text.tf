resource "harness_encrypted_text" "my_secret_text" {
  name              = "my_secret_text"
  value             = "foo"
  secret_manager_id = "6TLq_CboQ5CUtQyDev2yQg"

  usage_scope {
    application_filter_type = "ALL"
    environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
  }

  usage_scope {
    application_filter_type = "ALL"
    environment_filter_type = "PRODUCTION_ENVIRONMENTS"
  }

}

resource "harness_ssh_credential" "test" {
  name = "mysshcredential"
  ssh_authentication {
    port = 22
    username = "testuser"
    inline_ssh {
      passphrase_secret_id = harness_encrypted_text.my_secret_text.id
      ssh_key_file_id = "2WnPVgLGSZW6KbApZuxeaw"
    }

    # server_password {
    #   password_secret_id = harness_encrypted_text.my_secret_text.id
    # }

  }

  # kerberos_authentication {
  #   port = 25
  #   principal = "testuser"
  #   realm = "domain.com"
  #   tgt_generation_method {
  #     kerberos_password_id = harness_encrypted_text.my_secret_text.id
  #   }
  # }
}
