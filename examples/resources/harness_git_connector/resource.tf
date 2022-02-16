data "harness_secret_manager" "default" {
  default = true
}

resource "harness_encrypted_text" "example" {
  name              = "example-secret"
  value             = "foo"
  secret_manager_id = data.harness_secret_manager.default.id
}

resource "harness_git_connector" "example" {
  name                 = "example"
  url                  = "https://github.com/harness/terraform-provider-harness"
  branch               = "master"
  generate_webhook_url = true
  password_secret_id   = harness_encrypted_text.example.id
  url_type             = "REPO"
  username             = "someuser"
}
