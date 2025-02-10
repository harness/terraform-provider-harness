data "harness_platform_infra_variable_set" "test" {
  identifier = "identifier"
}

data "harness_platform_infra_variable_set" "testorg" {
  identifier              = "identifier"
  org_id                  = "someorg"
}

data "harness_platform_infra_variable_set" "testproj" {
  identifier              = "identifier"
  org_id                  = "someorg"
  project_id              = "someproj"
}
