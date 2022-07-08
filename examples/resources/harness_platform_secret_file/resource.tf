resource "harness_platform_secret_file" "example" {
  identifier                = "identifier"
  name                      = "name"
  description               = "test"
  tags                      = ["foo:bar"]
  file_path                 = "file_path"
  secret_manager_identifier = "harnessSecretManager"
}
