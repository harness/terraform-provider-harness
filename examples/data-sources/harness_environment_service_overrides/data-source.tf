data "harness_environment_service_overrides" "test" {
  identifier = harness_environment_service_overrides.test.id
  org_id     = harness_platform_organization.test.id
  project_id = harness_platform_project.test.id
  service_id = harness_platform_service.test.id
  env_id     = harness_platform_environment.test.id
}
