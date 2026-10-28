resource "harness_platform_connector_awscc" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  account_id  = "account_id"
  report_name = "report_name"
  report_type = "CUR2.0"
  region      = "us-east-1"
  s3_bucket   = "s3bucket"
  s3_prefix   = "cur"
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
