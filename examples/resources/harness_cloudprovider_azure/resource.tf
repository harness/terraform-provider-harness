data "harness_secret_manager" "default" {
  default = true
}

resource "harness_encrypted_text" "azure_key" {
  name              = "azure_key"
  value             = "<AZURE_KEY>"
  secret_manager_id = data.harness_secret_manager.default.id
}

resource "harness_cloudprovider_azure" "azure" {
  name      = "azure"
  client_id = "<AZURE_CLIENT_ID>"
  tenant_id = "<AZURE_TENANT_ID>"
  key       = harness_encrypted_text.azure_key.name
}
