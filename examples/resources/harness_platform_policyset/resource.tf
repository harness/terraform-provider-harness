resource "harness_platform_policyset" "test" {
  identifier = "harness_platform_policyset.test.identifier"
  name       = "harness_platform_policyset.test.name"
  action     = "onrun"
  type       = "pipeline"
  enabled    = true
  policies {
    identifier = "policy_identifier"
    severity   = "warning"
  }
}

## Policyset with multiple policies
resource "harness_platform_policyset" "test" {
  identifier = "harness_platform_policyset.test.identifier"
  name       = "harness_platform_policyset.test.name"
  action     = "onrun"
  type       = "pipeline"
  enabled    = true
  org_id     = "terraform_example_org"
  project_id = "terraform_test_project"
  policies {
    identifier = "policy_identifier1"
    severity   = "warning"
  }
  policies {
    identifier = "policy_identifier2"
    severity   = "warning"
  }
}
