data "harness_secret_manager" "default" {
  default = true
}

resource "harness_encrypted_text" "username" {
  name              = "k8s_username"
  value             = "<USERNAME>"
  secret_manager_id = data.harness_secret_manager.default.id
}

resource "harness_encrypted_text" "password" {
  name              = "k8s_password"
  value             = "<PASSWORD>"
  secret_manager_id = data.harness_secret_manager.default.id
}

resource "harness_cloudprovider_kubernetes" "example" {
  name            = "example-cluster"
  skip_validation = true

  authentication {
    username_password {
      master_url           = "https://localhost.com"
      username_secret_name = harness_encrypted_text.username.name
      password_secret_name = harness_encrypted_text.password.name
    }
  }
}
