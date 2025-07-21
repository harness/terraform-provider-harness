// Update image registry settings for service discovery
resource "harness_service_discovery_setting" "example" {
  org_identifier     = "sechaosworkshop"
  project_identifier = "se1"

  image_registry {
    account = "<account_name>"
    server  = "<registry_server>"
    secrets = ["<secret_name>"]
  }
}
