data "harness_secret_manager" "default" {
  default = true
}

resource "harness_encrypted_text" "aws_access_key" {
  name              = "aws_access_key"
  value             = "<ACCESS_KEY_ID>"
  secret_manager_id = data.harness_secret_manager.default.id
}

resource "harness_encrypted_text" "aws_secret_key" {
  name              = "aws_secret_key"
  value             = "<SECRET_KEY_ID>"
  secret_manager_id = data.harness_secret_manager.default.id
}

resource "harness_cloudprovider_aws" "aws" {
  name                          = "Example aws cloud provider"
  access_key_id_secret_name     = harness_encrypted_text.aws_access_key.name
  secret_access_key_secret_name = harness_encrypted_text.aws_secret_key.name
}
