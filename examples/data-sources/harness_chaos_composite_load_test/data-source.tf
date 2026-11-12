data "harness_chaos_composite_load_test" "checkout" {
  org_id     = "default"
  project_id = "demo"
  identifier = "checkout"
}
