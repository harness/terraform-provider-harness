// Data source to fetch service discovery setting by name
data "harness_service_discovery_setting" "example" {
  org_identifier     = "<org_identifier>"
  project_identifier = "<project_identifier>"
}
