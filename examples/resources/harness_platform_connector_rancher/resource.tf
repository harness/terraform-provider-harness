resource "harness_platform_connector_kubernetes" "bearer_token" {
  identifier  = "identifier"
  name        = "name"
  description = "description"
  tags        = ["foo:bar"]

  delegate_selectors = ["harness-delegate"]

  bearer_token {
    rancher_url  = "https://rancher.cluster.example"
    password_ref = "account.test_rancher_bearer_token"
  }
}
