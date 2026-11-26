data "harness_chaos_infrastructure_scan" "example" {
  org_id     = "default"
  project_id = "demo"
  identity   = "payments_infra_scan"
}
