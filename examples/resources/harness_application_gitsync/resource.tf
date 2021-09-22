data "harness_secret_manager" "default" {
  default = true
}

resource "harness_encrypted_text" "github_token" {
  name              = "github_token"
  value             = "<TOKEN>"
  secret_manager_id = data.harness_secret_manager.default.id
}

resource "harness_git_connector" "myrepo" {
  name                 = "myrepo"
  url                  = "https://github.com/someorg/myrepo"
  branch               = "main"
  generate_webhook_url = true
  username             = "someuser"
  password_secret_id   = harness_encrypted_text.github_token.id
  url_type             = "REPO"
}

resource "harness_application" "example" {
  name = "example-app"
}

resource "harness_application_gitsync" "example" {
  app_id       = harness_application.example.id
  connector_id = harness_git_connector.myrepo.id
  branch       = "main"
  enabled      = false
}
