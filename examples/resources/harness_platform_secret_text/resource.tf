resource "harness_platform_secret_text" "inline" {
  identifier  = "identifier"
  name        = "name"
  description = "example"
  tags        = ["foo:bar"]

  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "secret"
}

resource "harness_platform_secret_text" "reference" {
  identifier  = "identifier"
  name        = "name"
  description = "example"
  tags        = ["foo:bar"]

  secret_manager_identifier = "azureSecretManager"
  value_type                = "Reference"
  value                     = "secret"
}
