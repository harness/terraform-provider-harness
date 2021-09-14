data "harness_secret_manager" "default" {
  default = true
}

resource "harness_encrypted_text" "spot_token" {
  name              = "spot_token"
  secret_manager_id = data.harness_secret_manager.default.id
  value             = "<SPOT_TOKEN>"
}

resource "harness_cloudprovider_spot" "example" {
  name              = "example"
  account_id        = "<SPOT_ACCOUNT_ID>"
  token_secret_name = harness_encrypted_text.spot_token.name
}
