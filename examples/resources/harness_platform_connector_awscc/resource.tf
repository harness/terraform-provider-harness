resource "harness_platform_connector_awscc" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  account_id  = "account_id"
  report_name = "report_name"
  s3_bucket   = "s3bucket"
  features_enabled = [
    "OPTIMIZATION",
    "VISIBILITY",
    "BILLING",
  ]
  cross_account_access {
    role_arn    = "role_arn"
    external_id = "external_id"
  }
}
