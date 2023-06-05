resource "harness_platform_service_overrides_v2" "test" {
  identifier  = "identifier"
  org_id      = "orgIdentifier"
  project_id  = "projectIdentifier"
  service_id  = "serviceIdentifier"
  type        = "ENV_SERVICE_OVERRIDE"
}
