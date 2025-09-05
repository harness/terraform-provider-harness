data "harness_platform_policyset" "test" {
  identifier = "harness_platform_policyset.test.identifier"
  name       = "harness_platform_policyset.test.name"
  action     = "onrun"
  type       = "pipeline"
  enabled    = true
  policies_set {
    identifier = "always_run"
    severity   = "warning"
  }
}