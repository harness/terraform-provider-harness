resource "harness_platform_service_account" "example" {
  identifier  = "identifier"
  name        = "name"
  email       = "email@service.harness.io"
  description = "test"
  tags        = ["foo:bar"]
  account_id  = "account_id"
}
