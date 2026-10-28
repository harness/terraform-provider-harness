data "harness_chaos_pipeline_scan" "example" {
  org_id     = "default"
  project_id = "demo"
  identity   = "cart_deploy_scan"
}
