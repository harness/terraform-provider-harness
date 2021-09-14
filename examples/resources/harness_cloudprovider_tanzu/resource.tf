data "harness_secret_manager" "default" {
  default = true
}

resource "harness_encrypted_text" "tanzu_password" {
  name              = "tanzu_password"
  value             = "<PASSWORD>"
  secret_manager_id = data.harness_secret_manager.default.id
}

resource "harness_cloudprovider_tanzu" "example" {
  name                 = "example"
  endpoint             = "https://endpoint.com"
  skip_validation      = true
  username             = "<USERNAME>"
  password_secret_name = harness_encrypted_text.tanzu_password.name
}
