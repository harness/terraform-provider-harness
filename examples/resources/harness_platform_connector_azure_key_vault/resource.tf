resource "harness_platform_connector_azure_key_vault" "example" {
  identifier  = "identifier"
  name        = "name"
  description = "example"
  tags        = ["foo:bar"]

  client_id    = "client_id"
  secret_key   = "account.secret_key"
  tenant_id    = "tenant_id"
  vault_name   = "vault_name"
  subscription = "subscription"
  is_default   = false

  azure_environment_type = "AZURE"
}
