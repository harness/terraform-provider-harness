# Manual credentials example
resource "harness_platform_connector_azure_key_vault" "manual" {
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

# System-assigned Managed Identity example
resource "harness_platform_connector_azure_key_vault" "system_msi" {
  identifier  = "system_msi_example"
  name        = "system_msi_example"
  description = "Azure Key Vault using system-assigned managed identity"
  tags        = ["foo:bar"]

  vault_name   = "vault_name"
  subscription = "subscription"
  is_default   = false

  use_managed_identity        = true
  azure_managed_identity_type = "SystemAssignedManagedIdentity"

  delegate_selectors     = ["harness-delegate"]
  azure_environment_type = "AZURE"
}

# User-assigned Managed Identity example
resource "harness_platform_connector_azure_key_vault" "user_msi" {
  identifier  = "user_msi_example"
  name        = "user_msi_example"
  description = "Azure Key Vault using user-assigned managed identity"
  tags        = ["foo:bar"]

  vault_name   = "vault_name"
  subscription = "subscription"
  is_default   = false

  use_managed_identity        = true
  azure_managed_identity_type = "UserAssignedManagedIdentity"
  managed_client_id           = "client_id_of_managed_identity"

  delegate_selectors     = ["harness-delegate"]
  azure_environment_type = "AZURE"
}
