resource "harness_platform_secret_text" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "secret"
  lifecycle {
    ignore_changes = [
      value,
    ]
  }
}
