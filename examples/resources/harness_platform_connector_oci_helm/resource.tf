# Credential username_password
resource "harness_platform_connector_oci_helm" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "admin.azurecr.io"
  delegate_selectors = ["harness-delegate"]
  credentials {
    username     = "username"
    password_ref = "account.Secret_id"
  }
}

# Credential anonymous
resource "harness_platform_connector_oci_helm" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "admin.azurecr.io"
  delegate_selectors = ["harness-delegate"]
}
