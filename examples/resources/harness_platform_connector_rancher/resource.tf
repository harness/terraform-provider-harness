resource "harness_platform_connector_kubernetes" "bearer_token" {
  identifier  = "identifier"
  name        = "name"
  description = "description"
  tags        = ["foo:bar"]

  delegate_selectors = ["harness-delegate"]
  rancher_url        = "https://rancher.cluster.example"
  bearer_token {
    bearer_token_ref = "account.test_rancher_bearer_token"
  }
}
